package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	b byte
	x int32
}

//某些时候，获取对⻬信息对于内存⾃动分析是很有⽤的。
func main() {
	var d Data
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align()) // sizeof，以及最宽字段的对⻬模数。
	f, _ := t.FieldByName("b")
	fmt.Println(f.Type.FieldAlign()) // 字段对⻬。
}
