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

	fmt.Println("======================================")
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)           //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0

	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0

	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
	}

	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2) // []int{0, 0, 0}
	}

	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2)
	}

	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}

	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}
	fmt.Println("=================================")
	// 如果不特殊声明，go 的函数默认都是按值穿参，即通过函数传递的参数是值的副本，在函数内部对值修改不影响值的本身
	// 而slice,map较为特殊,无论是否用make初始化都 不用传递指针也会改变 但是固定长度的数组是不行的 仍旧是副本
	// 另外channel也是可以直接修改的
	s2 = make([]int, 3)
	fmt.Printf("%#v \n", s2) //[]int{0, 0, 0}
	modifySlice(s2)
	fmt.Printf("%#v \n", s2) //[]int{1, 0, 0}
	//而如果不是用slice初始化的
	s2 = []int{3, 3, 3}
	fmt.Printf("%#v \n", s2)
	modifySlice(s2)
	fmt.Printf("%#v \n", s2)
	// 改为数组而不是slice
	s6 := [3]int{3, 3, 3}
	fmt.Printf("%#v \n", s6)
	modifyArray(s6)
	fmt.Printf("%#v \n", s6)
	mt1 := map[string]interface{}{
		"1": 1,
	}
	fmt.Printf("%#v \n", mt1)
	modifyMap(mt1)
	fmt.Printf("%#v \n", mt1)
	var ct chan int
	if ct == nil {
		fmt.Printf("ct is nil --> %#v \n ", ct)
	}
	ct = make(chan int, 1)
	ct <- 1
	select {
	case <-ct:
		fmt.Println("ct <- 1")
		fmt.Printf("ct is nil --> %#v \n ", ct)
	}
}

func modifyMap(m map[string]interface{}) {
	m["1"] = 2
}

func modifySlice(s []int) {
	s[0] = 1
}

func modifyArray(s [3]int) {
	s[0] = 1
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
