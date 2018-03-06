// 创建一个文件存放数据,在同一时刻,可能会有多个Goroutine分别进行对此文件的写操作和读操作.
// 每一次写操作都应该向这个文件写入若干个字节的数据,作为一个独立的数据块存在,这意味着写操作之间不能彼此干扰,写入的内容之间也不能出现穿插和混淆的情况
// 每一次读操作都应该从这个文件中读取一个独立完整的数据块.它们读取的数据块不能重复,且需要按顺序读取.
// 例如: 第一个读操作读取了数据块1,第二个操作就应该读取数据块2,第三个读操作则应该读取数据块3,以此类推
// 对于这些读操作是否可以被同时执行,不做要求. 即使同时进行,也应该保持先后顺序.
package main

import (
	"fmt"
	"sync"
	"time"
	"os"
	"errors"
	"io"
)

//数据文件的接口类型
type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号 Read serial number
	Rsn() int64
	// 获取最后写入的数据块的序列号 Write serial number
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

//数据类型
type Data []byte

//数据文件的实现类型
type myDataFile struct {
	f *os.File	//文件
	rcond *sync.Cond
	fmutex sync.RWMutex //被用于文件的读写锁
	woffset int64 // 写操作需要用到的偏移量
	roffset int64 // 读操作需要用到的偏移量
	wmutex sync.Mutex // 写操作需要用到的互斥锁
	rmutex sync.Mutex // 读操作需要用到的互斥锁
	dataLen uint32 //读取和写入共同使用的数据块长度
}

//初始化DataFile类型值的函数,返回一个DataFile类型的值
func NewDataFile(path string, dataLen uint32) (DataFile, error){
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	//f,err := os.Create(path)
	if err != nil {
		fmt.Println("Fail to find", f, "cServer start Failed")
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f : f,
		dataLen:dataLen,
	}

	//df.fmutex.RLocker() 返回读写锁中的读锁 并与条件变量 rcond 进行绑定 条件变量不能与未锁定的写锁进行绑定以及未锁定的互斥锁 因为条件变量的wait会先释放锁 被唤醒后又会自动锁定锁
	df.rcond = sync.NewCond(df.fmutex.RLocker())

	return df, nil
}

//获取并更新读偏移量,根据读偏移量从文件中读取一块数据,把该数据块封装成一个Data类型值并将其作为结果值返回


func (df *myDataFile) Read() (rsn int64, d Data, err error){
	// 读取并更新读偏移量
	var offset int64
	// 读互斥锁定 普通互斥锁 保证读的有序性（偏移量）
	df.rmutex.Lock()
	offset = df.roffset
	// 更改偏移量, 当前偏移量+数据块长度
	df.roffset += int64(df.dataLen)
	// 读互斥解锁
	df.rmutex.Unlock()

	//读取一个数据块,最后读取的数据块序列号
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	//读写锁:读锁定
	df.fmutex.RLock()
	//必须解锁 wait 获得锁就会向下执行 这个时候必须释放锁 而wait没有获得锁的时候会释放锁 相当于一次锁定锁 一次释放锁
	defer df.fmutex.RUnlock()
	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			//fmt.Printf("df.f.ReadAt ERR: %+v\n",err)
			//由于进行写操作的Goroutine比进行读操作的Goroutine少,所以过不了多久读偏移量roffset的值就会大于写偏移量woffset的值
			// 也就是说,读操作很快就没有数据块可读了,这种情况会让df.f.ReadAt方法返回的第二个结果值为代表的非nil且会与io.EOF相等的值
			// 因此不应该把EOF看成错误的边界情况
			// 在读操作读完数据块,EOF时解锁读操作,并继续循环,尝试获取同一个数据块,直到获取成功为止.
			if err == io.EOF {
				//注意,如果在该for代码块被执行期间,一直让读写所fmutex处于读锁定状态,那么针对它的写操作将永远不会成功.
				//且相应的写Goroutine也会被一直阻塞.因为它们是互斥的.
				//所以 在每条return & continue 语句的前面加入一个针对该读写锁的读解锁操作
				//df.fmutex.RUnlock()
				// 和读写锁的读锁绑定后 则会自动挂起当前线程 释放读锁定
				df.rcond.Wait()
				//注意,出现EOF时可能是很多意外情况,如文件被删除,文件损坏等
				//这里可以考虑把逻辑提交给上层处理.
				//保留for循环 因为在唤醒后 仍然需要再次检查一遍是否仍然存在EOF的情况
				continue
			}else {
				//其他错误 直接返回
				//df.fmutex.RUnlock()
				return
			}
		}
		//不出现错误 则直接跳出循环
		break
	}
	d = bytes
	//必须去掉 否则会出现 panic: sync: RUnlock of unlocked RWMutex 因为放到上面的defer了
	//df.fmutex.RUnlock()
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error){
	//读取并更新写的偏移量
	var offset int64
	//互斥锁定
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	//互斥解锁
	df.wmutex.Unlock()

	//写入一个数据块,最后写入数据块的序号
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	//如果写入数据d太长 则进行截断
	if len(d) > int(df.dataLen){
		bytes = d[0:df.dataLen]
	}else{
		bytes = d
	}
	df.fmutex.Lock()
	//写入数据
	_, err = df.f.Write(bytes)
	if err != nil {
		return
	}
	//通知并唤醒某一个Goroutine进行读操作
	df.rcond.Signal()
	df.fmutex.Unlock()

	return
}

//计数值
func (df *myDataFile) Rsn() int64{
	//使用完全互斥锁
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

//计数值
func (df *myDataFile) Wsn() int64{
	//使用完全互斥锁
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func main(){
	//简单测试下结果
	//注意写入的字符长度必须和dataLen相同 否则可能出现无法完成全部读取 offset计算值和实际数据长度会不匹配
	var dataFile DataFile
	dataFile,_ = NewDataFile("./mutex_test", 5)

	var d=map[int]Data{
		1:[]byte("test1"),
		2:[]byte("test2"),
		3:[]byte("test3"),
	}

	//写入数据
	for i:= 1; i < 4; i++ {
		go func(i int){
			wsn,_ := dataFile.Write(d[i])
			fmt.Println("write i=", i,",wsn=",wsn, ",success.","data =",string(d[i]))
		}(i)
	}

	//读取数据
	for i:= 1; i < 4; i++ {
		go func(i int){
			rsn,d,_ := dataFile.Read()
			fmt.Println("Read i=", i,",rsn=",rsn,",data=",string(d), ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}