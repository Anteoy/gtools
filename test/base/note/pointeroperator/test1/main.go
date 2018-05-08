package main

import (
	"fmt"
	"unsafe"
)

// 注意：GC 把 uintptr 当成普通整数对象，它⽆法阻⽌ "关联" 对象被回收。
func main() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	n := 0
	n++
	n--
	p := uintptr(unsafe.Pointer(&d)) // 两次强制类型转换 普通指针 => unsafe.Pointer => uintptr

	p += unsafe.Offsetof(d.x) //指针地址从d.x的地址开始 偏移到第一个字段
	p2 := unsafe.Pointer(p)
	px := (*int)(p2) //p2转换为int型指针
	*px = 200        // d.x = 200
	//px = 200 这种方式会报编译错误

	fmt.Printf("%#v\n", d)
}
