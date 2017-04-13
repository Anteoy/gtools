package main

import (
	"fmt"
)

type Html []interface{}
//Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，T是普通类型。如果断言失败，ok为false，否则ok为true并且value为变量的值。来看个例子：
func main() {
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		if value, ok := element.(string); ok {
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}
	}
}