package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[0:1])
	fmt.Println(a[1:])
	fmt.Println(a[4:5])

	for _, v := range a {
		switch v {
		case 1:
			fmt.Println(1)
			continue
		case 2:
			fmt.Println(2)
			continue
		case 3:
			fmt.Println(3)
		default:
			fmt.Println("v.type err")
			break
		}
	}
	var wg *sync.WaitGroup
	wg = new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		defer func() {
			if aa := recover(); aa != nil {
				log.Println("[pannic recover]", aa)
			}
		}()
		panic("aa gua le main shi fou hui gua ")
	}(wg)
	go aa(wg)

	wg.Wait()
	fmt.Println("continue...")
	log.Fatalf("test err:%v\n", errors.New("ribenren"))
	fmt.Errorf("test:%s", errors.New("ribenren"))
	fmt.Printf("no")
}

func aa(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if aa := recover(); aa != nil {
			log.Println("[pannic recover]", aa)
		}
	}()
	panic("aa gua le main shi fou hui gua ")
}
