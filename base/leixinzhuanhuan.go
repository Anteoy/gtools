package main
//接口的转换遵循以下规则：

//普通类型向接口类型的转换是隐式的。

//接口类型向普通类型转换需要类型断言

//普通类型向接口类型转换的例子随处可见，例如
//"hello"作为string类型存储在interface{}类型的变量val中，[]byte{'a', 'b', 'c'}作为slice存储在interface{}类型的变量val中。这个过程是隐式的，是编译期确定的。
import (
	"fmt"
)

func main() {
	var val interface{} = "hello"
	fmt.Println(val)
	val = []byte{'a', 'b', 'c'}
	fmt.Println(val)
}