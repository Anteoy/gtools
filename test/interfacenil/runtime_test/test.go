package main

import (
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
		//defer wg.Done()
		defer func() {
			fmt.Println("wait!!! wait!!!")
			if aa := recover(); aa != nil {
				//这句话打印的位置不确定 非同步
				log.Println("[pannic recover]", aa)
			}
			//where
			wg.Done()
		}()
		//defer wg.Done()
		panic("aa gua le main shi fou hui gua ")
	}(wg)
	//go aa(wg)

	wg.Wait()
	fmt.Println("continue...")
	//log.Fatalf("test err:%v\n", errors.New("ribenren"))
	err := fmt.Errorf("test11111:%s", "ribenren")
	fmt.Printf("..... %+v\n", err)
	//这个代码必须在panic代码之前编码,否则无效
	defer func() {
		fmt.Println("zhu1!!! panic1!!!")
		if aa := recover(); aa != nil {
			fmt.Println("zhu2!!! panic2!!!")
		}
	}()
	//continue panic 这个也不是顺序执行的 打印的日志为随机的
	panic("zhu xie cheng panic!!!")
	//panic recover之后不再继续执行
	fmt.Println("no")
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
