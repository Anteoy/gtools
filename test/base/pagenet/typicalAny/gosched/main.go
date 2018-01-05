package main

import (
	"fmt"
	"runtime"
	"sync"
)

const N = 26

func main() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)
	fmt.Printf("%c\n", 'A'+1)
	fmt.Printf("%c\n", 'A'+2)
	fmt.Printf("%c\n", 'A'+25)
	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			// A
			//有序和无序 debug顺序会不一定 yields the processor
			runtime.Gosched()
			fmt.Printf("%d ===>%c\n", i,'a'+i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d ===>%c\n",i, 'A'+i)
		}(i)
	}
	wg.Wait()
}