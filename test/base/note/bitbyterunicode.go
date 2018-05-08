package main

import "fmt"

func main() {
	//default UTF-8 (unicode impl)
	s := "Go编程"
	// 2 + 3 + 3
	fmt.Println(len(s))
	// 和上面相同 汉字在UTF-8 (unicode 的实现表示方法之一,不需要每个unicode字符都占用32位,4个字节byte)
	fmt.Println(len([]byte(s)))
	// rune 为 int32, byte 为 uint8 ,所以1个 rune能包括所有的一个unicode字符
	fmt.Println(len([]rune(s)))

}
