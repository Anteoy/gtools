package main

import (
	"fmt"
	"math"
)

type aliaint int

func main(){
	const x  = 100
	const y byte = x
	// 这种状态的常量不能直接赋值给byte 编译器认为是两种类型
	const xx  int= 100
	//const yy  byte= xx

	a,b,c := 100,0144,0x64
	fmt.Println(a,b,c)

	fmt.Println(math.MinInt8,math.MaxInt8)
	// 不能这么用 只有内置别名byte uint8 和 rune int32可以 其他类型需要自己转换
	//var  ax int  = 6
	//var _ aliaint = ax
	//var a byte = 0x11
	//var _ uint8 = a

	var d byte = 0x11
	var bb uint8 = d
	var _ uint8 = d+ bb

}
