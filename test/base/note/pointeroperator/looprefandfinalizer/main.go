package main

import (
	"fmt"
	"runtime"
	"time"
)

type Data struct {
	d [1024 * 100]byte
	o *Data
}

// 指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露。
func test() {
	var a, b Data
	//循环引⽤
	a.o = &b
	b.o = &a
	//垃圾回收器能正确处理 "指针循环引⽤" ，但⽆法确定 Finalizer 依赖次序，也就⽆法调⽤
	//Finalizer 函数，这会导致⺫标对象⽆法变成不可达状态，其所占⽤内存⽆法被回收。
	//runtime.SetFinalizer 在一个对象 obj 被从内存移除前执行一些特殊操作
	runtime.SetFinalizer(&a, func(d *Data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { fmt.Printf("b %p final.\n", d) })
}

// go build -gcflags "-N -l" && GODEBUG="gctrace=1" ./looprefandfinalizer
func main() {
	for {
		test()
		time.Sleep(time.Millisecond)
	}
}
