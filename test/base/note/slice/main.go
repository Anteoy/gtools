package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5}
	data1 := [6]int{0, 1, 2, 3, 4, 5}
	//相同类型 [6]int 数组
	fmt.Printf("%T\n", data)
	fmt.Printf("%T\n", data1)
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	fmt.Println(s)
	data[5] = 30
	fmt.Println(data)
	fmt.Println("======")

	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使⽤索引号。8:100 不足8的全部为0,然后再第9赋值一个100
	fmt.Println(s1, len(s1), cap(s1))
	s4 := []int{0, 1, 2, 3, 8}
	fmt.Println(s4, len(s4), cap(s4))
	s2 := make([]int, 6, 8) // 使⽤ make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}
