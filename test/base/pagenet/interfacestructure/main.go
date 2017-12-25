package main

import (
	"fmt"
)

type People interface {
	Show()
}
type People1 interface {
	//Show()
}

//People and People1 both not == nil because :
//type eface struct {      //空接口 is nil
//	_type *_type         //类型信息
//	data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
//}
//type iface struct {      //带有方法的接口 is not nil
//	tab  *itab           //存储type信息还有结构实现方法的集合 is not nil
//	data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
//}

var in interface{}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	if in == nil {
		fmt.Println("CCCCCC")
	}else{
		fmt.Println("DDDDDD")
	}
}