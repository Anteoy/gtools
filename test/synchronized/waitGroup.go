// WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.sina.com/",
		"http://www.baidu.com/",
		"http://www.anteoy.site/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("wait over")
}

// 从执行结果可看出：
// 1、取三个网址信息的时候，结果显示顺序与for循环的顺序没有必然关系。
// 2、三个goroutine全部执行完成后，wg.Wait()才停止等待，继续执行并打印出over字符。
