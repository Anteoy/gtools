package main

import (
	"errors"
	"fmt"
	"time"
	// "os"
)

func main() {
	e := errors.New("error")
	fmt.Println(e)
	// panic(e) //（3） defer 不会执行 没有声明就已经panic了
	// os.Exit(1) //（4） defer 不会执行 主动调用 os.Exit(int) 退出进程时，defer 将不再被执行。
	defer fmt.Println("defer")
	// go func() { panic(e) }() //（1） 会导致 defer 不会执行 在发生 panic 的（主）协程中，如果没有一个 defer 调用 recover()进行恢复，则会在执行完最后一个已声明的 defer 后，引发整个进程崩溃；
	// panic(e) //（2） defer 会执行
	time.Sleep(1e9)
	fmt.Println("over.")
	// os.Exit(1) //（5） defer 不会执行  主动调用 os.Exit(int) 退出进程时，defer 将不再被执行。
}