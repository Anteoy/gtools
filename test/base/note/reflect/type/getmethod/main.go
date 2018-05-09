package main

import (
	"fmt"
	"reflect"
)

type User struct {
}
type Admin struct {
	User
}

func (*User) ToString() {}

func (Admin) test() {}
func main() {
	var u Admin
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}
	fmt.Println("--- value interface ---")
	// 什么输出都没有
	methods(reflect.TypeOf(u))
	fmt.Println("--- pointer interface ---")
	// 只有指针方法 ToString
	methods(reflect.TypeOf(&u))
}
