package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

//除 struct，其他复合类型 array、 slice、 map 取值⽰例。
func main() {
	v := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v.Len(); i < n; i++ {
		fmt.Println(v.Index(i).Int())
	}
	fmt.Println("---------------------------")
	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}

	u := User{}
	vv := reflect.ValueOf(u)

	//f := vv.FieldByName("age")
	f := vv.FieldByName("a")
	//Value 某些⽅法没有遵循 "comma ok" 模式，⽽是返回 ZeroValue，因此需
	//要⽤ IsValid 判断⼀下是否可⽤。
	fmt.Println(f.Kind(), f.IsValid())
}
