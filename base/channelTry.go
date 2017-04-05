package main

import "fmt"
//17 -5 12 或者 -5 17 12
func sum1(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println("sum1向channel发送total: "+string(total))
	c <- total  // send total to c
	fmt.Println("sum1测试是否阻塞！！！")
}

func sum2(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	fmt.Println("sum2向channel发送total: " + string(total))
	// 修改后 会称为死锁
	//fatal error: all goroutines are asleep - deadlock!
	if false{
		c <- total  // send total to c
	}
	fmt.Println("sum2测试是否阻塞！！！")
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum1(a[:len(a)/2], c)
	go sum2(a[len(a)/2:], c)
	fmt.Println("main测试是否阻塞！！！start wait for receive!!!")
	//缺省情况下，发送和接收会一直阻塞着，知道另一方准备好。这种方式可以用来在gororutine中进行同步，而不必使用显示的锁或者条件变量
	//一直等待计算结果发送到channel中.
	x, y := <-c, <-c  // receive from c
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("main测试是否阻塞！！！end wait for receive!!!")
	fmt.Println(x, y, x + y)

}
