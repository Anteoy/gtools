package main

/*
 #include <stdio.h>
 #include <stdlib.h>
 void test(char *s) {
 printf("%s\n", s);
 }
 char* cstr() {
 return "abcde";
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//  go build main.go
// ./main
//Hello, World!
//abcde
//ab
//[97 98]
func main() {
	s := "Hello, World!"
	cs := C.CString(s)               // 该函数在 C heap 分配内存，需要调⽤ free 释放。
	defer C.free(unsafe.Pointer(cs)) // #include <stdlib.h>
	C.test(cs)

	cs = C.cstr()
	fmt.Println(C.GoString(cs))
	// 参数2 同下
	fmt.Println(C.GoStringN(cs, 2))
	// 第二个参数2 表示会获取数组的前两个数 2的时候打印 [97 98] ; 1的时候打印[97]
	fmt.Println(C.GoBytes(unsafe.Pointer(cs), 2))
}
