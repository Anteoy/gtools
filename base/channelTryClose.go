package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		//如果把close(c)注释掉，程序会一直阻塞在for …… range那一行 并不会finish完成程序并退出
		//close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
