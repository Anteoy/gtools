package main

import (
	"sync"
	"fmt"
	"time"
)

func main(){
	var mutex sync.Mutex
	fmt.Println("LOCK G0")
	mutex.Lock()
	fmt.Println("LOCKED G0")
	for i:=0;i<3;i++{
		go func(i int){
			fmt.Printf("LOCK G(%d)\n",i)
			mutex.Lock()
			fmt.Printf("LOCKED G(%d)\n",i)
		}(i)
	}
	time.Sleep(time.Second*2)
	mutex.Unlock()
	fmt.Println("UNLOCKED")
	time.Sleep(time.Second*10)

}
