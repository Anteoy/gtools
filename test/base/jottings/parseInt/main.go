package main

import (
	"fmt"
	"strconv"
)

//ref: https://stackoverflow.com/questions/21491488/what-is-the-difference-between-int-and-int64-in-go
func main() {
	//i, err := strconv.ParseInt("-128", 10, 8)
	//i, err := strconv.ParseInt("128", 10, 8)
	i, err := strconv.ParseInt("127", 10, 8)
	fmt.Println(i, err)
}
