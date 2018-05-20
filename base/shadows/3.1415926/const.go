package  main

import (
	"fmt"
	"reflect"
)

const Pi float64 = 3.14159265358979323846
const zero = 0.0
const notzero = 1.3

const (
	size int64 = 1024
	eof = -1
)
// nil 必须指定类型 并且不能赋值给string类型 如果是1之类的就不用
//var x = nil
//var x string = nil
var xx interface{} = nil

type Aa interface{}
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"

const (
	// errors.New 不是一个常量 error 不是一个有效的常量类型
	//ERR_ELEM_EXIST error = errors.New("element already exists")
	//ERR_ELEM_NT_EXIST error = errors.New("element not exists")
)
func main(){
	//print float64
	fmt.Println(reflect.TypeOf(zero).Name())
	//print float64
	fmt.Println(reflect.TypeOf(notzero).Name())
	//print int
	fmt.Println(reflect.TypeOf(eof).Name())
	a := [2]int{1,2}
	fmt.Println(cap(a))
	b := make(chan<- int,21)
	fmt.Println(cap(b),len(b))

	m := make(map[int]int,2)
	m[0] = 0
	m[1] = 1
	fmt.Println(len(m))
	fmt.Println("--------------")
	mtest(m)
	fmt.Println(m)

	fmt.Println("------------")
	x := 1
	{
		x := 2
		fmt.Printf("%p,%d\n",&x,x)
	}
	fmt.Printf("%p,%d\n",&x,x)
	fmt.Println("++++++++++++++++")
	xx := []string{"a", "b", "c"}
	for v := range xx {
		fmt.Print(v)
	}
	fmt.Println("----------------")
	str:= "hello"
	str1 := []byte(str)
	str1[0] = 'x'
	fmt.Println(string(str1))
}

func mtest(m map[int]int) {
	m[2] = 2
	m[1] = 3
	delete(m,0)
}