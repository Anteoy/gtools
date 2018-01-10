package main

import (
	"fmt"
)

// type 1
//a defer1: 1
//a defer2: 2
//a return: 0

//type 2
//a defer1: 1 0xc42000a368
//a defer2: 2 0xc42000a368
//a return: 0xc42000a360 0
func main() {
	r := a()
	fmt.Println("a return:", &r,r) // 打印结果为 a return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i,&i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i,&i) // 打印结果为 a defer1: 1
	}()
	//return 的时候 新建了一个返回值接收的变量 相当于一次赋值 而最开始就赋值了 并且赋值的是值而不是指针 所以返回的是0 而不是改变后的2
	return i
}