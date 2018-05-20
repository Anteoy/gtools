package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	wg := sync.WaitGroup{}
	for i:=0;i<2;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("start: %d\n",i)
			defer func(i int) {
				fmt.Printf("defer:%d\n",i)
			}(i)
			if i == 1 {
				time.Sleep(time.Second*10)
				//其他携程的defer不会执行 这里只打印了defer: 1
				panic("2 panic")
			}else if i == 0 {
				time.Sleep(time.Second *20)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}

//func (a unsafe.Pointer) Test () {
//
//}
