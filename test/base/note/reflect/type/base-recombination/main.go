package main

import (
	"fmt"
	"reflect"
)

var (
	Int    reflect.Type = reflect.TypeOf(0)
	String reflect.Type = reflect.TypeOf("")
)

// 可从基本类型获取所对应复合类型。
func main() {
	var c reflect.Type
	c = reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)
	m := reflect.MapOf(String, Int)
	fmt.Println(m)
	s := reflect.SliceOf(Int)
	fmt.Println(s)
	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)
	//与之对应，⽅法 Elem 可返回复合类型的基类型。
	var t1 reflect.Type
	t1 = reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t1)
}
