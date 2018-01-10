package main

import (
	"fmt"
)

//b defer1: 1
//b defer2: 2
//b return: 2
//非匿名的返回值 i int ,一开始就声明了i，return 虽然先赋值（重要）了返回值i，但是因为i是可变的，就像指针一样，所以执行完defer后 i的指针值已变为2
//defer、return、返回值三者的执行顺序应该是：return最先给返回值赋值；接着 defer 开始执行一些收尾工作；最后 RET 指令携带返回值退出函数。
func main() {
	fmt.Println("b return:", b()) // 打印结果为 b return: 2
}
//这里非匿名理解为不会新建返回值变量了 直接返回i 而i是可变的引用变量 所以也会改变
func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}