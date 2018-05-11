package main

import (
	"fmt"
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	// not ptr .Elem()
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
}

type test struct {
	Name string
	Sex  int
}

func main() {
	structName := "test"

	s, ok := newStruct(structName)
	if !ok {
		return
	}
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println("==============")

	ss := s.(test)
	ss.Name = "ss"
	ss.Sex = 1
	fmt.Println(ss, reflect.TypeOf(ss))

}
