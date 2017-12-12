package main

import (
	"time"
	"fmt"
)

func main(){
	fmt.Println("time for 2 s")
	limiter := time.Tick(time.Millisecond * 2000)
	<- limiter
	fmt.Println("end1")
	time.Sleep(10*time.Second)
	fmt.Println("end")
}
