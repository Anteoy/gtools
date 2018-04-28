package main

/*
#cgo  CFLAGS:  -I  ./include
#cgo  LDFLAGS:  -L ./lib  -lhello
#include "hello.h"
*/
import "C"

func main() {
	C.hello(C.CString("call C hello func"))
}
