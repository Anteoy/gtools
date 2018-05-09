package main

/*
 #include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	//void* 对应golang里面unsafe.Pointer 分配了32字节 malloc申请内存 可能不会把内存中的值清掉
	m := unsafe.Pointer(C.malloc(4 * 8))
	defer C.free(m) // 注释释放内存。
	//1个int 32bit 4byte 4个int的数组就是4*4=16字节 还有16个字节去哪儿了? 原来这里是4 我改为了8 18 超过申请的内存了
	p := (*[18]int)(m) // 转换为数组指针。
	fmt.Println(p)
	for i := 0; i < 18; i++ {
		p[i] = i + 100
	}
	fmt.Println(p)
}
