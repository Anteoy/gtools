package main

import "fmt"

func main(){
	add(1,2)
	add(1,3,7)
	add([]int{1,3,5}...)
}

func add(args ...int){
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
}
