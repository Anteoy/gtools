package main

/*
 #cgo CFLAGS: -I.
 #cgo LDFLAGS: -L. -ltest
 #include "test.h"
*/
import "C"

func main() {
	println(C.sum(10, 20))
}
