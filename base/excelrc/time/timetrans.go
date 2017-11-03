package main

import (
	"fmt"
	"time"
)

func main() {
	//timeFormat := "2006-01-02 15:04"
	timeFormat := "20060102 15:04"
	time2, _ := time.Parse(timeFormat, "20161128 19:36") //洛杉矶时间
	fmt.Printf("%v\n", time2)
	fmt.Printf("%v\n", time2.Unix())
}
