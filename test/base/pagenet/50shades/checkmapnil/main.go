package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	x := map[string]string{"one":"a","two":"","three":"c"}
	if v := x["two"]; v == "" { //incorrect
		fmt.Println("no entry")
	}

	if _,ok := x["two"]; !ok {
		fmt.Println("no entry2")
	}

	xx := "text"
	fmt.Println(xx[0]) //print 116
	fmt.Println(string(xx[0])) //print t
	fmt.Printf("%T\n",xx[0]) //prints uint8


	//字符串的值不需要是UTF8的文本。它们可以包含任意的字节。只有在string literal使用时，字符串才会是UTF8。即使之后它们可以使用转义序列来包含其他的数据。
	//为了知道字符串是否是UTF8，你可以使用“unicode/utf8”包中的ValidString()函数。
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1)) //prints: true
	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2)) //prints: false

	data := "♥"
	fmt.Println(len(data)) //prints: 3

	//RuneCountInString()函数并不返回字符的数量，因为单个字符可能占用多个rune。
	//rune 能操作 任何字符
	//byte 不支持中文的操作
	//Go语言中byte和rune实质上就是uint8和int32类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	data11 := "é"
	fmt.Println(len(data11))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data11)) //prints: 2
}
