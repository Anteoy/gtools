package main

import (
	"fmt"
	"time"
)
//发送者将不会被阻塞，除非消息正在被接收者处理。根据你运行代码的机器的不同，接收者的goroutine可能会或者不会有足够的时间，在发送者继续执行前处理消息。
func main(){
	ch := make(chan int)

	go func() {
		for m:=range ch{
			fmt.Println(m)
		}
	}()

	ch <- 11
	ch <- 12
	ch <- 14
	time1 := time.Tick(2*time.Second)
	<-time1
	for i:=0;i<10;i++{
		ch <- i
	}
	time11 := time.Tick(2*time.Second)
	<-time11
	ch <- 13
}
