package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := []int64{1, 2, 3}
	b := []string{"1a", "2b", "3c"}
	fmt.Println(fmt.Sprintf("%v\n", a))
	fmt.Println(fmt.Sprintf("%v\n", b))

	bb, _ := json.Marshal(a)
	aa, _ := json.Marshal(b)
	fmt.Println(string(aa))
	fmt.Println(string(bb))
}
