package main

import (
	"runtime"
	"sync"
	"fmt"
)
//A:均为输出10，B:从0~9输出(顺序不定)。
// 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。 故go func执行时，i的值始终是10。
func main(){
	//WU LUN SHI DUO SHAO PROC
	//DOU HUI CHU XIAN
	//fatal error: all goroutines are asleep - deadlock!
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<10;i++{
		//ZUI ZHONG I = 10
		go func(){
			fmt.Println("A: ",i)
			wg.Done()
		}()
	}
	for i:=0;i<10;i++{
		go func(i int){
			fmt.Println("b :",i)
		}(i)
	}
	wg.Wait();
}
