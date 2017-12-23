package main

import "fmt"

func init(){

}
func main(){
	if false == false {
		fmt.Println(111)
	}
	a := "hello"

	var s []int
	s = append(s,1)
	ss := []int{1,2,3}
	//6 must have end ,
	//_ = []int{1,2,3,
	//4,5,6,
	//}
	_ = []int{1,2,3,
	4,5,6}
	s = append(s,ss...)
	//error bu neng assign
	//a[0] = 'x'
	//i := 1
	//meiyou
	 //++i
	 //must have len
	 //s := make([]int)
	fmt.Println(a[0],string(a[0]))
	fmt.Println(a[1],string(a[1]))
}