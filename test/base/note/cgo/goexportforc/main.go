// 导出 Go 函数给 C 调⽤，须使⽤ "//export" 标记。建议在独⽴头⽂件中声明函数原型，
//避免 "duplicate symbol" 错误。
package main

import "fmt"

/*
 #cgo  CFLAGS:  -I  ./

 #include "test.h"
*/
import "C"

//export hello
func hello() {
	fmt.Println("Hello, World!\n")
}

func main() {
	C.test()
}
