package main

import(
	"time"
	"fmt"
)

func main() {
	a := 1
	go func(){
		a = 2
	}()
	a = 3
	//fangdao zhe er shi 2 ,fang dao xia mian shi 3 (yi ban qing kuang)
	time.Sleep(2 * time.Second)
	fmt.Println("a is ", a)

}