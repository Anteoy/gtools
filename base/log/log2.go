package main

import (
	"runtime"
	"fmt"
)

func main() {
	funcName, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println("Func Name=" + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s    line=%d\n", file, line)
	}
}