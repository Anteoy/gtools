package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

//TODO
func main() {
	f, err := os.Open("file")
	defer f.Close()
	defer fmt.Println("11111111111")
	defer func() {
		fmt.Println("2222222222")
	}()
	if err != nil {
		return
	}
	//correct
	//defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}