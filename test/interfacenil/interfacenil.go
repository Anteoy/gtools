package main

import "fmt"

type AA struct {
}

func main() {
	var a *AA
	//a = new(AA)
	fmt.Printf("a : %+v\n", a)
	var b interface{} = a
	fmt.Println(a == nil, b == nil, b == a)
}
