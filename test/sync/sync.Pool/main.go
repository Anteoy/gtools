package main

import (
	"runtime/debug"
	"sync/atomic"
	"sync"
	"fmt"
	"runtime"
)

//临时对象池
//自动伸缩，高效，安全
//可能会被垃圾回收器回收 随时可能被其他P偷走
//就算垃圾回收了 也可能没回收到其中的部分 比如下面的垃圾回收后仍然能获取值 但是如果把new参数置为nil 则获取到的也只有nil了
//Get获取的值不是确定的
//Set也不确定会到哪个P的本地池
func main(){
	//禁用gc 并在main结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newfunc := func () interface{}{
		return atomic.AddInt32(&count,1)
	}
	pool:=sync.Pool{New:newfunc}
	//new 字段值的作用
	v1 := pool.Get()
	fmt.Println("v1:",v1)

	//临时对象池的存取
	pool.Put(newfunc())
	pool.Put(newfunc())
	pool.Put(newfunc())
	v2 := pool.Get()
	fmt.Println("v2:",v2)

	//垃圾回收器对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()
	//就算垃圾回收了 也可能没回收到其中的部分 比如下面的垃圾回收后仍然能获取值 但是如果把new参数置为nil 则获取到的也只有nil了
	v3 := pool.Get()
	fmt.Println("v3:",v3)
	//nil后只能获取到nil了
	pool.New = nil
	v4 := pool.Get()
	fmt.Println("v4:",v4)
}
