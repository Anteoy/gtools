package main

import "fmt"

func test() {
	x, y := 10, 20
	defer func(i int) {
		println("defer:", i, y) // y 闭包引⽤
	}(x) // x 被复制 (在程序开始 初始化时即已复制了原始值)
	x += 10
	y += 100
	println("x =", x, "y =", y)
}

func main() {
	test()
	test2()
	test3()
}

// 延迟调⽤中引发的错误，可被后续延迟调⽤捕获，但仅最后⼀个错误可被捕获。
func test2() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

// 捕获函数 recover 只有在延迟调⽤内直接调⽤才会终⽌错误，否则总是返回 nil。任何未
// 捕获的错误都会沿调⽤堆栈向外传递。
func test3() {
	defer recover()              // ⽆效！
	defer fmt.Println(recover()) // ⽆效！
	defer func() {
		func() {
			println("defer inner")
			recover() // ⽆效！
		}()
	}()
	panic("test panic")
}
