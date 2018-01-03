package main

import (
	"fmt"
	"os"
)

type Slice []int
func NewSlice() Slice {
	return make(Slice, 0)
}
func (s* Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
// all 13122
func main() {
	s := NewSlice()
	//132
	defer s.Add(1).Add(2)
	//312
	defer func() {
		s.Add(1).Add(2)
		//can not
		//delete(s,1)
	}()
	s.Add(3)

	//TODO  maybi yichang pinic?
	file, err := os.Open("test.go")
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
}
//type in interface{}
//can not build error
//func (in in)a(){
//
//}
