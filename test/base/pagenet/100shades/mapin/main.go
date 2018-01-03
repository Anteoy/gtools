package main

import "fmt"

func main(){
	m1 := map[int]int{
		1:1,
		2:2,
		3:3,
	}
	a(m1)

	for k,v := range m1{
		fmt.Println( fmt.Sprintf("%d:%d\n",k,v))
	}

	fmt.Println("=========")
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}

}

func a(m1 map[int]int) {
	m1[4] = 4
	m1[3] = 33
}