package main

import (
	"fmt"
	"time"
)
//defer 延迟执行的只有函数体 参数不是 defer声明时会先计算确定参数的值
func main() {
	//时间在main前面
	defer P(time.Now())
	time.Sleep(5e9)
	fmt.Println("main ", time.Now())
}

func P(t time.Time) {
	fmt.Println("defer", t)
	fmt.Println("P    ", time.Now())
}