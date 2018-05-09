package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}
type Admin struct {
	User
	title string
}

func main() {
	var u Admin
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("title")
	fmt.Println(f.Name)
	f, _ = t.FieldByName("User") // 访问嵌⼊字段。
	fmt.Println(f.Name)
	f, _ = t.FieldByName("Username") // 直接访问嵌⼊字段成员，会⾃动深度查找。
	fmt.Println(f.Name)
	f = t.FieldByIndex([]int{0, 1}) // Admin[0] -> User[1] -> age
	fmt.Println(f.Name)
}
