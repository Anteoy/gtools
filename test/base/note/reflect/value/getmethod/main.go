package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}
func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}
func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i))
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}
}

//可获取⽅法参数、返回值类型等信息。
func main() {
	d := new(Data)
	t := reflect.TypeOf(d)
	test, _ := t.MethodByName("Test")
	info(test)
	sum, _ := t.MethodByName("Sum")
	info(sum)

	fmt.Println("-----------------------")
	//动态调⽤⽅法很简单，按 In 列表准备好所需参数即可 (不包括 receiver)。
	v := reflect.ValueOf(d)
	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)
		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}
	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
	fmt.Println("-----------------------")
	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(3),
	})

	//如改⽤ CallSlice，只需将变参打包成 slice 即可。
	fmt.Println("-----------------------")
	m := v.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}), // 将变参打包成 slice。
	}
	out := m.CallSlice(in)
	for _, v := range out {
		fmt.Println(v.Interface())
	}
}
