package main

import (
	"time"
	"fmt"
	"strconv"
)

func main(){
	now := time.Now()
	hh, _ := time.ParseDuration("2h")
	//hh, _ := time.ParseDuration("1m")
	hh1 := now.Add(hh)
	int64aa := hh1.UnixNano()
	fmt.Println(strconv.FormatInt(int64aa,10))
}