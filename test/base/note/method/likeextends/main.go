package main

import "fmt"

type User struct {
	id   int
	name string
}

//通过匿名字段，可获得和继承类似的复⽤能⼒。依据编译器查找次序，只需在外层定义同
//名⽅法，就可以实现 "override"。
type Manager struct {
	//匿名字段
	User
}

//类似多态重写
func (m *Manager) ToString() string {
	return fmt.Sprintf("Manger: %p, %#v", m, m)
}

//类似继承
func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %#v", self, self)
}
func main() {
	m := Manager{User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}
