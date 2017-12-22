package main

import "fmt"

type myInt int


func main(){
	var i int = 1
	//类型别名为新的类型 即myInt和int不是同一类型 不能直接赋值
	//var j myInt = i
	//但可以进行强制类型转换
	var j myInt = myInt(i)
	//不是interface类新 不能断言
	//var j myInt = i.(myInt)
	fmt.Println(j)
}
