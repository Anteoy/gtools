package main

import "fmt"

type istruct struct {
	A int
}

type Message struct {
	Data  map[string]string `json:"data,omitempty"`
	Data1 []int
	Data2 [2]int
	Data3 istruct
}

func main() {
	a := &Message{}
	fmt.Println(a.Data)
	//add
	a.Data = make(map[string]string, 10)
	//panic: assignment to entry in nil map
	a.Data["test"] = "test1"
	//panic: runtime error: index out of range
	//a.Data1[0] = 1
	//0,0
	fmt.Println(len(a.Data1), cap(a.Data1))
	a.Data1 = []int{}
	//0,0
	fmt.Println(len(a.Data1), cap(a.Data1))
	//ok 10 10
	a.Data1 = make([]int, 10, 10)
	a.Data1[0] = 1

	a.Data2[0] = 1

	a.Data3.A = 1
	fmt.Println(a)

	//panic: runtime error: index out of range
	//bb := []int{}
	//bb[0] = 1
	//fmt.Println(bb)

	b := new(Message)
	fmt.Println(b.Data)
	//add
	b.Data = map[string]string{}
	//panic: assignment to entry in nil map
	b.Data["test"] = "test1"
	fmt.Println(b)

}
