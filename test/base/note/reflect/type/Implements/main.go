package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) String() string {
	return ""
}
func main() {
	var d *Data
	t := reflect.TypeOf(d)
	// 没法直接获取接⼝类型，好在接⼝本⾝是个 struct，创建
	// ⼀个空指针对象，这样传递给 TypeOf 转换成 interface{}
	// 时就有类型信息了。。
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	//⽅法 Implements 判断是否实现了某个具体接⼝
	// 为啥不是 t.Implements(fmt.Stringer)，完全可以由编译器⽣成。
	fmt.Println(t.Implements(it))
}
