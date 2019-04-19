package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	//c <- 1
	//c <- 2
	//c <- 3
	//close(c)

	//for {
	//i, isClose := <-c
	//if !isClose {
	//	fmt.Println("channel closed!")
	//	break
	//}
	//fmt.Println(i)
	//}

	//for i := range c {
	//	fmt.Println(i)
	//}

	select {
	case i := <-c:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}

	fmt.Println("exec over")
}
