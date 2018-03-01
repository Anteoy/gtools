package main

import "fmt"

type data struct {
	name string
}
func main() {
	//m := map[string]*data {"x":{"one"}}
	//panic: runtime error: invalid memory address or nil pointer dereference
	//	[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x401103]
	//m["z"].name = "what?" //???
	//cannot assign to struct field m["x"].name in map
	//m["x"].name = "two" //error

	s := []data{{"x"}}
	s[0].name = "two" //ok
	fmt.Println(s)    //prints: [{two}]
}