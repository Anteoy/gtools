package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Username string
	age      int
}

//将对象转换为接⼝，会发⽣复制⾏为。该复制品只读，⽆法被修改。所以要通过接⼝改变
//目标对象状态，必须是 pointer-interface。
//就算是指针，我们依然没法将这个存储在 data 的指针指向其他对象，只能透过它修改目
//标对象。因为目标对象并没有被复制，被复制的只是指针。
func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	p1 := reflect.ValueOf(&u)
	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p1.CanSet(), p1.Elem().FieldByName("Username").CanSet())

	//⾮导出字段⽆法直接修改，可改⽤指针操作。
	p := reflect.ValueOf(&u).Elem()
	p.FieldByName("Username").SetString("Tom")
	f := p.FieldByName("age")
	fmt.Println(f.CanSet())
	// 判断是否能获取地址。
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}
	// 注意 p 是 Value 类型，需要还原成接⼝才能转型。
	fmt.Println(u, p.Interface().(User))
}
