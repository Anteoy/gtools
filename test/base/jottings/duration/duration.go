package main

import "fmt"

//类型别名和类型定义最大的区别在于：类型别名和原类型是相同的，而类型定义和原类型是不同的两个类型。
type Duration int64
//type D = int
//type I int
func main(){
	var a Duration = 1234
	var b int64 = 123
	//无效的操作符 类型不匹配
	//if a == b {
	//
	//}
	//同上
	//b = a
	b = int64(a)
	a = Duration(b)
	fmt.Println(a)
	fmt.Println("+++")
	fmt.Println(b)
}
