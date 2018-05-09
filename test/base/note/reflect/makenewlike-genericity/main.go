package main

import (
	"fmt"
	"reflect"
)

var (
	Int    reflect.Type = reflect.TypeOf(0)
	String reflect.Type = reflect.TypeOf("")
)

//原理并不复杂。
//1. 核⼼是提供⼀个 swap 函数，其中利⽤ reflect.MakeSlice ⽣成最终 slice 对象，
//因此需要传⼊ element type、 len、 cap 参数。
//2. 接下来，利⽤ reflect.MakeFunc 函数⽣成 swap value，并修改函数变量指向，以达到调
//⽤ swap 的目地。
//3. 当调⽤具体类型的函数变量时，实际内部调⽤的是 swap，相关代码会⾃动转换参
//数列表，并将返回结果还原成具体类型返回值。
//如此，在共享算法的前提下，⽆须⽤ interface{}，⽆须做类型转换，颇有泛型的效果。
func Make(T reflect.Type, fptr interface{}) {
	// 实际创建 slice 的包装函数。
	swap := func(in []reflect.Value) []reflect.Value {
		// --- 省略算法内容 这里可以添加自己的逻辑 本例子中直接使用reflect.MakeSlice()方法 返回一个根据传入类型而创建的指定类型的slice --- //
		// 返回和类型匹配的 slice 对象。
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T), // slice type
				int(in[0].Int()),   // len
				int(in[1].Int()),   // cap
			),
		}
	}

	// 传⼊的是函数变量指针，我们要将变量指向 swap 函数。
	fn := reflect.ValueOf(fptr).Elem()

	// 获取函数指针类型，⽣成所需 swap function value。
	v := reflect.MakeFunc(fn.Type(), swap)

	// 修改函数指针实际指向，也就是 swap。
	fn.Set(v)
}
func main() {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string
	// ⽤相同算法，⽣成不同类型创建函数。
	Make(Int, &makeints)
	Make(String, &makestrings)
	// 按实际类型使⽤。
	//创建一个[]int slice,len 5 cap 10
	x := makeints(5, 10)
	fmt.Printf("%#v\n", x)
	//创建一个[]string slice,len 3 cap 10
	s := makestrings(3, 10)
	fmt.Printf("%#v\n", s)
}
