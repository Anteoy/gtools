package main

import "fmt"

type A struct {
	BB []*B
}

type B struct {
}

func main(){
	c := &A{}
	fmt.Printf("%+v,%+v,%d,%d\n",c,c.BB,len(c.BB),cap(c.BB))
	//can not use c.BB[0] not exist
	d := make([]*B,1,1)
	fmt.Printf("%+v",d[0])
	d[0] = nil
}
