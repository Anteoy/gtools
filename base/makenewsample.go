package main

import "fmt"

type foo struct {
	A int
	B string
}

//make 只能创建 slice,map 和 channel;分配了一个有初始值（非零）的T类型（我这里打印的为空，个人理解为空值）返回值 当slice cap无法满足需要，则在初始大小cap（即第三个参数，这里为6），成倍增长
//new 分配了零值填充的T类型的内存空间，并返回其地址（指针） 当slice cap无法满足需要，则以16为基数，成倍增长
func main() {
	a := make([]int, 2, 6)
	b := new([]int)
	c := new(foo)
	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b", b, *b, len(*b), cap(*b))
	for i := 0; i < 10; i++ { //i < 10,i <20
		a = append(a, 2)
		*b = append(*b, 2)
	}
	fmt.Println("=======================================")
	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("b:", b, *b, len(*b), cap(*b))
	fmt.Println("=======================================")
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	//同样由int的零值0填充了
	fmt.Println(c.A)
	//同样由string的零值""空字符串填充了
	fmt.Printf("%+v\n", c)
	fmt.Println(len(c.B))
}

// i < 10
// a: [0 0] 2 6
// b &[] [] 0 0
// =======================================
// a: [0 0 2 2 2 2 2 2 2 2 2 2] 12 12
// b: &[2 2 2 2 2 2 2 2 2 2] [2 2 2 2 2 2 2 2 2 2] 10 16
// i < 20
// a: [0 0] 2 6
// b &[] [] 0 0
// =======================================
// a: [0 0 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2] 22 24
// b: &[2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2] [2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2] 20 32
