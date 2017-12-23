package main

import "fmt"

func main(){
	//var a bool = 1 no
	//var a bool = bool(1)
	//var i = 10 ok
	var m map[string]int
	m["a"]= 1
 	var bt b = &a1{}
	fmt.Println(bt)
}

type a1 struct {

}

func (p *a1)aa(){

}
func (p *a1)bb(){

}

type b1 struct{

}
func (p *b1)aa(){

}

type a interface{
	bb()
	aa()
}

type b interface{
	aa()
}