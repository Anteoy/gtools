package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 8,9
	s := data[8:]
	// 0,1,2,3,4
	s2 := data[:5]
	copy(s2, s) // dst:s2, src:s
	//8,9,2,3,4
	fmt.Println(s2)
	//8 9 2 3 4 5 6 7 8 9
	fmt.Println(data)
}
