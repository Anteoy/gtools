package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("main lock")
	mutex.Lock()
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("lock i=%d\n", i)
			mutex.Lock()
			fmt.Printf("lock ok i=%d\n", i)
		}(i)
	}
	runtime.Gosched()
	time.Sleep(time.Second)
	fmt.Println("unlock by main")
	mutex.Unlock()
	fmt.Println("unlocked by main")
	time.Sleep(time.Second)

}
