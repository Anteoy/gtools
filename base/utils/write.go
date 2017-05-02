package main

import (
	"bufio" //缓存IO
	"fmt"
	"io"
	"io/ioutil" //io 工具包
	"os"
	"time"
)

const (
	TEST_WRITING_LINE    = 10000000
	TEST_WRITING_CONTENT = "Go是Google开发的一种编译型，可平行化，并具有垃圾回收功能的编程语言。"
)

func writeString(writer func(string), f *os.File) {
	writer(TEST_WRITING_CONTENT)
	if f != nil {
		f.Close()
	}
}

/**
  from: http://www.isharey.com/?p=143
*/
func main() {
	var t time.Time
	t = time.Now()
	writeString(writeToFile1("testfile1.txt"))
	fmt.Printf("使用 io.WriteString 写入文件:%v\n", time.Now().Sub(t))
	t = time.Now()
	// writeString(writeToFile2("testfile2.txt"))
	// fmt.Printf("使用 ioutil.WriteFile 写入文件:%v\n", time.Now().Sub(t))
	// t = time.Now()
	writeString(writeToFile3("testfile3.txt"))
	fmt.Printf("使用 File(Write,WriteString) 写入文件:%v\n", time.Now().Sub(t))
	t = time.Now()
	writeString(writeToFile4("testfile4.txt"))
	fmt.Printf("使用 bufio.NewWriter 写入文件:%v\n", time.Now().Sub(t))
}

func writeToFile1(filename string) (func(string), *os.File) {
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	var f *os.File
	var err error
	f, err = os.Create(filename) //创建文件
	if err != nil {
		panic(err)
	}

	return func(content string) {
		_, err := io.WriteString(f, content) //写入文件(字符串)
		if err != nil {
			panic(err)
		}
	}, f
}
func writeToFile2(filename string) (func(string), *os.File) {
	/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	return func(content string) {
		var bytes = []byte(content)
		err := ioutil.WriteFile(filename, bytes, 0666) //写入文件(字节数组)
		if err != nil {
			panic(err)
		}
	}, nil

}
func writeToFile3(filename string) (func(string), *os.File) {
	/*****************************  第三种方式:  使用 File.Write/File.WriteString 写入文件 ***********************************************/
	f, err := os.Create(filename) //创建文件
	if err != nil {
		panic(err)
	}
	return func(content string) {
		var bytes = []byte(content)
		_, err := f.Write(bytes) //写入文件(字节数组)
		if err != nil {
			panic(err)
		}
	}, f
}
func writeToFile4(filename string) (func(string), *os.File) {
	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	f, err := os.Create(filename) //创建文件
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	return func(content string) {
		_, err = w.WriteString(content) //写入文件(字节数组)
		if err != nil {
			panic(err)
		}
		w.Flush()
	}, f
}
