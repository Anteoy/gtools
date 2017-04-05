package main

import (
	"fmt"
	"reflect"
)

func main(){
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	//如果修改值，则reflect 必须传入xx的指针 即&xx 否则会发生panic
	var xx float64 = 3.4
	vv := reflect.ValueOf(&xx).Elem()
	vv.SetFloat(7.1)
	fmt.Println(vv)
	fmt.Println(xx)
}