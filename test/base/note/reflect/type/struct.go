package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}
type Admin struct {
	User
	title string
}

func main() {
	var u Admin
	t := reflect.TypeOf(u)
	//t := reflect.TypeOf(u.User)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	//如果是指针，应该先使⽤ Elem ⽅法获取目标类型，指针本⾝是没有字段成员的。
	u1 := new(Admin)
	t1 := reflect.TypeOf(u1)
	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}
	for i, n := 0, t1.NumField(); i < n; i++ {
		f := t1.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}
