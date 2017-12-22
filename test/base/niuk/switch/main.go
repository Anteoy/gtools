package main

import "fmt"

func main(){

	//只支持一个赋值初始化语句
	if a, b := 21, 3; a > b {
		fmt.Println("a>b ? true")
	}

	//for i, j := 1, 10; i < j; i,j=i+1,j+1 {  //可死循环
	//	fmt.Println(i)
	//}
	//for {//可死循环
	//	fmt.Println("aa")
	//}
	//由于Go没有逗号表达式，而++和--是语句而不是表达式
	//for i, j := 1, 10; i < j; i++,j++ {
	//	fmt.Println(i)
	//}

	//switch可以没有表达式，在 Case 中使用布尔表达式，这样形如 if-else 不只可以跟字符串和数字的表达式,其他可识别的类型也可以
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}


}
