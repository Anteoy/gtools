// file name: test.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int8
}

func testMarshal() []byte {
	user := User{
		Name: "Tab",
		Age:  18,
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func testUnmarshal(data []byte) {
	var user User
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func testRead() []byte {
	fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:n]))
	return data[:n]
}

func testWrite(data []byte) {
	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//var data []byte
	//data = testMarshal()
	//fmt.Println(string(data))
	//testWrite(data)
	//data = testRead()
	//testUnmarshal(data)

	var test [3]string = [3]string{"t1"}
	//test := make([]string, 1)
	add3(test)
	fmt.Printf("%+v\n  ", test)
}

func add1(s *[]string) {
	*s = append(*s, "tetst")
}

func add2(s []string) {
	s[0] = "test"
}

func add3(s [3]string) {
	s[0] = "test"
}
