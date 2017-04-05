package main

import (
	"fmt"
)

func test(c, quit chan int) {
	x, y := 0, 1
	//var val interface{} = x
	//fmt.Println(val.(string))
	// fmt.Println(val.(int))
	var i int8 = 0
	for {
		select {
		//向c发送值x
		case c <- x:
			//panic fmt.Println("向c已发送x:%d" + val.(string))
			fmt.Printf("向c第%d次已发送x: %d\n", i, x)
			//fmt.Println(string(x)) 乱码
			//fmt.Println("向c已发送x: ",  strconv.Itoa(x))
			x, y = y, x+y
			fmt.Println("计算后x的值为%d\n",x)
			fmt.Println("计算后y的值为%d\n",y)
			i++
		//接收quit值
		case <-quit:
		//打印quit
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			//从c接收值并打印
			fmt.Printf("从channel c中第%d次接收到值：%d\n", i,<-c)
		}
		quit <- 0
	}()
	test(c, quit)

	//test
	//x,y := 7,11
	////后下面的不一样
	//x, y = y, x+y
	////x = y
	////y = x + y
	//fmt.Println("计算后x的值为%d\n",x)
	//fmt.Println("计算后y的值为%d\n",y)

}
