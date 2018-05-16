package main

import (
	"fmt"
	"sync"
)

func main() {
	//全部都是打印10 或者可能打印有其他值了(比如我有一次打印出了7和3各一个)
	//这个主要看goroutine执行时机 在匿名函数的闭包对象实际是一个指针,这个指针就是i的指针,只不过这个时候
	//i是多少,即ii是多少是无法确定的
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Printf("i (%p) = %d\n", &i, i)
		wg.Add(1)
		go func() {
			fmt.Printf("ii (%p) = %d\n", &i, i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("+++++++++++++++++++++++++")
	//正常打印 j和jj同样其实是一个指针地址
	for j := 0; j < 10; j++ {
		fmt.Printf("j (%p) = %d\n", &j, j)
		func() {
			fmt.Printf("jj (%p) = %d\n", &j, j)
		}()
	}

	//闭包复制的是原对象指针，这就很容易解释延迟引⽤现象。
	fmt.Println("+++++++++++++++++++++++++")
	f := test()
	f()
}

func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
