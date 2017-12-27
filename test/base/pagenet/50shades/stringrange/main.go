package main

import "fmt"

func main() {
	data := "A\xfe\x02\xff\x04"
	//byte的索引,bu shi rune de suo ying。它不是当前“字符”的索引，这与其他语言不同。注意真实的字符可能会由多个rune表示。如果你需要处理字符，确保你使用了“norm”包（golang.org/x/text/unicode/norm）。
	for index,v := range data {
		fmt.Printf("%#x \n",v)
		fmt.Printf("%d \n",index)
	}
	//prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

	//string变量的for range语句将会尝试把数据翻译为UTF8文本。对于它无法理解的任何byte序列，它将返回0xfffd runes（即unicode替换字符），而不是真实的数据。如果你任意（非UTF8文本）的数据保存在string变量中，确保把它们转换为byte slice，以得到所有保存的数据。
	fmt.Println("---------")
	for _,v := range []byte(data) {
		fmt.Printf("%#x ",v)
		fmt.Printf("%s \n",string(v))
	}
	fmt.Println("---------")
	for _,v := range []rune(data) {
		fmt.Printf("%#x ",v)
		fmt.Printf("%s \n",string(v))
	}
	//prints: 0x41 0xfe 0x2 0xff 0x4 (good)
}
