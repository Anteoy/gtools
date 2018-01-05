package main

import (
	"fmt"
	"time"
)

func main(){
	//ch := make(chan int)
	//for i := 0; i < 3; i++ {
	//	go func(idx int) {
	//		ch <- (idx + 1) * 2
	//	}(i)
	//}
	//
	////get the first result
	//fmt.Println(<-ch)
	//close(ch) //not ok (you still have other senders) panic: send on closed channel
	////do other work
	//time.Sleep(2 * time.Second)

	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2: fmt.Println(idx,"sent result")
			case <- done: fmt.Println(idx,"exiting")
			}
		}(i)
	}
	//get first result
	fmt.Println("result:",<-ch)
	close(done)
	//do other work
	time.Sleep(3 * time.Second)
}
