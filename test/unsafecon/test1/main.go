package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int
	B int
}

func main() {
	foo := &Foo{1, 2}
	fmt.Println(foo)

	base := uintptr(unsafe.Pointer(foo))
	offset := unsafe.Offsetof(foo.A)

	ptr := unsafe.Pointer(base + offset)
	*(*int)(ptr) = 3

	fmt.Println(foo)
}
