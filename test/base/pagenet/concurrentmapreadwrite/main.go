package main

import (
	"sync"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	//ua.Lock()
	//defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
//fatal error: concurrent map read and map write
func main(){
	ug := &UserAges{}
	ug.ages = make(map[string]int)
	fmt.Println(ug.ages)
	wg := sync.WaitGroup{}
	wg.Add(100)
	//10 ge ti xian bu chu lai
	for i:=0;i<100;i++{
		go func(i int){
			ug.Add("i"+string(i),i)
			wg.Done()
		}(i)
		//must there
		go func(i int){
			fmt.Println(ug.Get("i"+string(i)))
		}(i)
	}
	wg.Wait()
	//for i:=0;i<100;i++{
	//	go func(i int){
	//		ug.Get("i"+string(i))
	//	}(i)
	//}

}