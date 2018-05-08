package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	// 原来的slice不会改变 而是返回了新的slice
	fmt.Println(s, s2)

	//data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//// 原来的array会改变 s2为新的struct slice,同样指向原来数组
	//s := data[:3]
	//s2 := append(s, 100, 200) // 添加多个值。
	//fmt.Println(data)
	//fmt.Println(s)
	//fmt.Println(s2)
	//data[0] = 100
	////改变data 则s2也改变了
	//fmt.Println(s2)
}
