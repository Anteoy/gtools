package main

import "fmt"

//defer hou jin xian chu
//panic... wei zhi bu que ding
func main(){
	call()
}

func call(){
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")

	panic("panic...")
}
