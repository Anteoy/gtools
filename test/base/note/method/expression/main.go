package main

import "fmt"

type User struct {
	id   int
	name string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self User) Test2() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	u := User{1, "Tom"}
	u.Test()
	mValue := u.Test
	// 这一句和上面的效果是相同的
	//mValue := (&u).Test
	// 内存地址都是相同的哦 这里虽然没有显示传递指针,单实际隐示传递了指针
	mValue() // 隐式传递 receiver   =>  method value 绑定user实例
	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver => method expression 传入user实例

	fmt.Println("======")
	//需要注意，method value 会复制 receiver
	u1 := User{1, "Tom"}
	mValue1 := u1.Test2 // ⽴即复制 receiver，因为不是指针类型，不受后续修改影响。
	u1.id, u1.name = 2, "Jack"
	u1.Test()
	mValue1()
	//在汇编层⾯，method value 和闭包的实现⽅式相同，实际返回 FuncVal 类型对象。
}
