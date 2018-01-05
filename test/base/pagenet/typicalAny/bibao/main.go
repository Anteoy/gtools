package main


import (
	"sync"
	"fmt"
	"strconv"
)

const N = 10
//ke neng wei 1 keneng wei 2
func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	for _,v := range m {
		fmt.Println(":"+strconv.Itoa(v) )
	}
}
