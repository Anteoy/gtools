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

//Value 和 Type 使⽤⽅法类似，包括使⽤ Elem 获取指针⺫标对象。 Type 获取类型相关 Value获取值相关(也可以获取Type相关,不知道是否可以获取所有type能获取的东西,我估计不行)
func main() {
	u := &Admin{User{"Jack", 23}, "NT"}
	v := reflect.ValueOf(u).Elem()
	fmt.Println(v.FieldByName("title").String())   // ⽤转换⽅法获取字段值
	fmt.Println(v.FieldByName("age").Int())        // 直接访问嵌⼊字段成员
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int()) // ⽤多级序号访问嵌⼊字段成员

	fmt.Println(v.FieldByName("Username").Interface())
	//除具体的 Int、 String 等转换⽅法，还可返回 interface{}。只是⾮导出字段⽆法使⽤，需
	//⽤ CanInterface 判断⼀下。
	if v.FieldByName("age").CanInterface() {
		//否则直接使用会pannic
		fmt.Println(v.FieldByName("age").Interface())
	} else {
		fmt.Println("can not age")
	}
	//当然，转换成具体类型不会引发 panic。
	fmt.Println(v.FieldByName("age").Int())

	fmt.Println("--------------------")
	var p *int
	var x interface{} = p
	fmt.Println(x == nil)
	vv := reflect.ValueOf(p)
	fmt.Println(vv.Kind(), vv.IsNil())
}
