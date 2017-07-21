package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func MapToStruct(obj interface{}, objMap map[string]interface{}) error {
	for name, value := range objMap {
		structValue := reflect.ValueOf(obj).Elem()
		structFieldValue := structValue.FieldByName(name)

		if !structFieldValue.CanSet() {
			return fmt.Errorf("struct 字段 %s 不能set,请检查!!!", name)
		}

		if !structFieldValue.IsValid() {
			return fmt.Errorf("struct中找不到 %s 的对应字段列,请检查!!!", name)
		}

		structFieldType := structFieldValue.Type()
		val := reflect.ValueOf(value)
		if structFieldType != val.Type() {
			return fmt.Errorf("map struct 同名匹配列的类型不匹配,请检查!!! struct: %s,map value: %s\n", structFieldType, val.Type())
		}

		structFieldValue.Set(val)
	}

	return nil
}

type MyStruct struct {
	Name string
	Age  int
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = 23

	result := &MyStruct{}
	//err := result.FillStruct(myData)
	err := MapToStruct(result, myData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
