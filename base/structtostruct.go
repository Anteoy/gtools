package main

import "fmt"
//基本上只有两个完全一样的struct(就是名字不同，成员定义、顺序、包括tag都得一样)才可以直接转换，转换直接t2 := Test1(t1)这样就可以(t1是Test1类型)
//前提是Test1, Test2符合转换条件，哪怕tag定义不一样都不行的，比如下面这样都不能转换的，编译会报错。
type Test1 struct {
	Age int      /*`json:"age"`*/
	Name string
}
type Test2 struct {
	Age int
	Name string
}
func main(){
	aa := Test1{Age:11}
	//bb := &Test2{Age:11}

	cc := Test2(aa)
	fmt.Println(cc.Age)
}
