package main
import "fmt"

import (
"reflect"
)
//反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”。
//注：这里反射类型指 reflect.Type 和 reflect.Value。
//反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”。
//反射第三定律：如果要修改“反射类型对象”，其值必须是“可写的”（settable）
//https://segmentfault.com/a/1190000006190038
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	//“可写性”是反射类型变量的一个属性，但不是所有的反射类型变量都拥有这个属性。
	//我们可以通过 CanSet 方法检查一个 reflect.Value 类型变量的“可写性”。
	//“可写性”有些类似于寻址能力，但是更严格。它是反射类型变量的一种属性，赋予该变量修改底层存储数据的能力。“可写性”最终是由一个事实决定的：反射对象是否存储了原始值。举个代码例子：
	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//这里我们传递给 reflect.ValueOf 函数的是变量 x 的一个拷贝，而非 x 本身。想象一下，如果下面这行代码能够成功执行：
	//v.SetFloat(7.1)
	//答案是：如果这行代码能够成功执行，它不会更新 x ，虽然看起来变量 v 是根据 x 创建的。相反，它会更新 x 存在于 反射对象 v 内部的一个拷贝，而变量 x 本身完全不受影响。这会造成迷惑，并且没有任何意义，所以是不合法的。“可写性”就是为了避免这个问题而设计的。
	//这看起来很诡异，事实上并非如此，而且类似的情况很常见。考虑下面这行代码：
	//f(x)
	//上面的代码中，我们把变量 x 的一个拷贝传递给函数，因此不期望它会改变 x 的值。如果期望函数 f 能够修改变量 x，我们必须传递 x 的地址（即指向 x 的指针）给函数 f，如下：
	//f(&x)
	//你应该很熟悉这行代码，反射的工作机制是一样的。如果你想通过反射修改变量 x，就咬吧想要修改的变量的指针传递给 反射库
	vv := reflect.ValueOf(x)
	fmt.Println("settability of v:", vv.CanSet())
}
