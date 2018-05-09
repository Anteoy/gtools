package main

import (
	"fmt"
	"reflect"
	"strings"
)

func makeAdd(T interface{}) {

	//通用算法函数体模板 MakeFunc第二个参数的固定格式
	add := func(args []reflect.Value) (results []reflect.Value) {

		if len(args) == 0 {

			return nil

		}

		var r reflect.Value

		switch args[0].Kind() {

		case reflect.Int:

			n := 0

			for _, a := range args {

				n += int(a.Int())

			}

			r = reflect.ValueOf(n)

		case reflect.String:

			ss := make([]string, 0, len(args))

			for _, s := range args {

				ss = append(ss, s.String())

			}

			r = reflect.ValueOf(strings.Join(ss, "<join>"))

		}

		results = append(results, r)

		return

	}

	fn := reflect.ValueOf(T).Elem()

	v := reflect.MakeFunc(fn.Type(), add) //把原始函数变量的类型和通用算法函数存到同一个Value中 这里返回的v是一个有函数签名和实现的function

	fn.Set(v) //把原始函数指针变量指向v，这样它就获得了函数体

}

func main() {

	//定义函数变量，未定义函数体
	var intAdd func(x, y int) int

	var strAdd func(a, b string) string

	makeAdd(&intAdd)

	makeAdd(&strAdd)

	fmt.Println(intAdd(12, 23)) //35

	fmt.Println(strAdd("hello,", "world!")) //hello, world!

}
