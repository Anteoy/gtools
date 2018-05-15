package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	//这个不行 因为指针都指向一个临时地址 是一个指针
	list2 := make([]*Foo, len(list))
	//for i, value := range list {
	//	list2[i] = &value
	//}

	// 这个可以
	//list2 := make([]Foo, len(list))
	//for i, value := range list {
	//	list2[i] = value
	//}

	//var value Foo
	//for i := 0; i < len(list); i++ {
	//	value = list[i]
	//	list2[i] = &value
	//}

	//正确的例子 解决方案1 用i来访问,解决方案2,list1初始化为*Foo指针
	for i, _ := range list {
		list2[i] = &list[i]
	}

	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2])

}
