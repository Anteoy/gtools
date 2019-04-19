package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		fmt.Println(111111111111111)
	}
}
