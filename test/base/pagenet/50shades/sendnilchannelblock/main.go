package main
import (
	"fmt"
	"time"
)
//fatal error: all goroutines are asleep - deadlock!
//在一个nil的channel上发送和接收操作会被永久阻塞。这个行为有详细的文档解释，但它对于新的Go开发者而言是个惊喜。
//suo yi ye hui pannic
func main() {
	var ch chan int
	var aa int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			//both can not ..
			//ch <- (idx + 1) * 2
			select{
			case ch <- aa:

			}

		}(i)
	}
	//get first result
	//fmt.Println("result:",<-ch)
	//do other work
	time.Sleep(2 * time.Second)

	//这个行为可以在select声明中用于动态开启和关闭case代码块的方法。
	inch := make(chan int)
	outch := make(chan int)
	go func() {
		var in <- chan int = inch
		var out chan <- int
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <- in:
				out = outch
				in = nil
			}
		}
	}()
	time.Sleep(3*time.Second)
	go func() {
		for r := range outch {
			fmt.Println("result:",r)
		}
	}()
	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
