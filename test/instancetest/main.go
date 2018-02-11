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


	fmt.Println("=================")
	arr := make([]*B,1)
	fmt.Println("len,cap: ",len(arr),cap(arr))
	b := &B{}
	arr = append(arr,b)
	fmt.Println("=================")
	fmt.Println("len,cap: ",len(arr),cap(arr))
	//会增加一个nil在里面
	fmt.Printf("%+v\n",arr)
	fmt.Println("+++++++++++++++")
	arr1 := make([]*B,1)
	arr1[0] = b
	fmt.Println("len,cap: ",len(arr1),cap(arr1))
	fmt.Printf("%+v\n",arr1)
	fmt.Println("+++++++++++++++")
	arr2 := make([]*B,0)
	fmt.Println("len,cap: ",len(arr2),cap(arr2))
	arr2 = append(arr2,b)
	fmt.Println("len,cap: ",len(arr2),cap(arr2))
	fmt.Printf("%+v\n",arr2)
}
