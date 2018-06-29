package main

import "fmt"

type Message struct {
	Data map[string]string `json:"data,omitempty"`
}

func main() {
	a := &Message{}
	fmt.Println(a.Data)
	//add
	a.Data = make(map[string]string, 10)
	//panic: assignment to entry in nil map
	a.Data["test"] = "test1"
	fmt.Println(a)

	b := new(Message)
	fmt.Println(b.Data)
	//add
	b.Data = map[string]string{}
	//panic: assignment to entry in nil map
	b.Data["test"] = "test1"
	fmt.Println(b)
}
