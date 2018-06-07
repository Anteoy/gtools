package main

func main(){
	const x  = 100
	const y byte = x
	// 这种状态的常量不能直接赋值给byte 编译器认为是两种类型
	const xx  int= 100
	//const yy  byte= xx
}
