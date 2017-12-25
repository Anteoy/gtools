package main

import "fmt"

type stu struct{
	Name string
	Age int
}

//range 循环，会重用地址，也就是说，for _, stu := range stus 中的 stu 总是在同一个地址
//考察内容为range的行为，range其实是在做值拷贝，stu变量的地址始终没有改变，最终stu地址指向的内容也就是最后一个数据。
func main(){
	m := make(map[string]*stu)
	stus := []stu{
		{Name:"r1",Age:1},
		{Name:"r2",Age:2},
		{Name:"r4",Age:4},
		{Name:"r3",Age:3},
	}
	for _,stu := range stus {
		m[stu.Name] = &stu
	}

	for _,stu := range m {
		fmt.Printf("%s:%d\n",stu.Name,stu.Age)
	}


}
