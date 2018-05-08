package main

import "fmt"

func main() {
	mm := make(map[int]int, 1000)
	fmt.Println(len(mm)) //0
	//fmt.Println(cap(mm)) //nocap
	fmt.Println(mm)

	type user struct{ name string }
	m := map[int]user{ // 当 map 因扩张⽽重新哈希时，各键值项存储位置都会发⽣改变。 因此，map
		1: {"user1"}, // 被设计成 not addressable。 类似 m[1].name 这种期望透过原 value
	} // 指针修改成员的⾏为⾃然会被禁⽌。
	//m[1].name = "Tom" // Error: cannot assign to m[1].name

	//正确做法是完整替换 value 或使⽤指针。
	u := m[1]
	u.name = "Tom"
	m[1] = u // 替换 value。

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。

	fmt.Println("=======")

	// 可以在迭代时安全删除键值。但如果期间有新增操作，那么就不知道会有什么意外了。 产生未知结果
	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			//新增
			m[k+k] = "x"
			//删除
			delete(m, k)
		}
		fmt.Println(m)
	}
}
