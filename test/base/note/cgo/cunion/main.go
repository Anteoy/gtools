package main

/*
 #include <stdlib.h>
 union Data {
	 char x;
	 int y;
 };
 union Data* test() {
 	union Data* p = malloc(sizeof(union Data));
	p->x = 100;
	return p;
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var d *C.union_Data = C.test()
	defer C.free(unsafe.Pointer(d))
	fmt.Println(d)
}
