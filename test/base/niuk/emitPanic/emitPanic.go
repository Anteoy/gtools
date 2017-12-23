package main

import (
	"os"
	"fmt"
)

func main(){
	file,err := os.Open("test.go")
	defer file.Close()
	if err != nil {
		fmt.Println("open file faild:",err)
		return
	}
}
