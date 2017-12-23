package main

import (
	"reflect"
	"fmt"
)
const a,b float32 = 0,3
const aa,bb,c = 3,4,"foo"

const (
	//err
	//err1 error = errors.New("abc")
	q =3
	e int64 = 87
)
func main(){
	a := 0.0
	const zero = 0.0
	fmt.Println(reflect.TypeOf(a),reflect.ValueOf(a).Kind())
	fmt.Println(reflect.TypeOf(zero),reflect.ValueOf(zero).Kind())

	aa := []int{1,2,3}
	fmt.Println(cap(aa))
	fmt.Println(len(aa))
	bb := [2]int{1,2}
	fmt.Println(cap(bb))
	fmt.Println(len(bb))
	cc := make(chan int,4)
	fmt.Println(cap(cc))
	fmt.Println(len(cc))
	dd := make(map[string]string,2)
	//map no cap ;have len
	fmt.Println(len(dd))

}
