package main

import "fmt"

type Integer int

func (a Integer) Add(b Integer) Integer{
	return  a + b
}
//*Integer Integer 不是同一类型 不能 a + b,并且传递的参数类型也强制进去区分Integer和*Integer,只能a + *b(b为 *Integer)这里的*b为解除引用
//func (a Integer) Add(b *Integer) Integer{
//	return  a + *b
//}

//IDE 不报错,但是使用go build会报错,IDE目前不能智能提示这个错误
//cannot call pointer method on i.(Integer)
//此处类型断言 i.(Integer)会报错
//func (a *Integer) Add(b Integer) Integer{
//	return  *a + b
//}

func main(){
	var a Integer = 1
	var b Integer = 2
	var i interface{} = a
	sum := i.(Integer).Add(b)
	fmt.Println(sum)
}
