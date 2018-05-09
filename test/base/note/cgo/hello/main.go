package main

/*
 #include <stdio.h>
 #include <stdlib.h>
 void hello() {
 	printf("Hello, World!\n");
 }
*/
import "C"

// 注意 必须使用命令行进行编译运行  如 go run main.go  使用IDE等工具 不能保证有gcc g++编译器参与 不能保证程序正常运行和调试
func main() {
	C.hello()
}
