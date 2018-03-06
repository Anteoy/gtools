package main

import (
	"sync"
	"fmt"
	"time"
)

//once.Do 只会执行一次 如数据库的连接池的初始化任务
func main(){
	var num int
	sign := make(chan bool)
	var once sync.Once
	f := func(i int) func(){
		return func() {
			num =(num + i*2)
			sign <- true
		}
	}
	for i:=0;i<3;i++{
		fi := f(i+1)
		go once.Do(fi)
	}
	for j:=0;j<3;j++{
		select{
		case <-sign:
			fmt.Println("shou dao yi ge sign")
		case <-time.After(100*time.Millisecond):
			fmt.Println("time out!!!")
		}
	}
	fmt.Println("num =",num)
}
