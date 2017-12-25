package main

import "fmt"

func main(){
	s := make([]int,6)
	//s := make([]int,6,10)
	// 6 ,6 ru guo you di san ge  ze cap wei di san ge fou ze dou shi di er ge
	// cap shi zhi jie x2 fan bei
	//6,10
	fmt.Println(len(s),cap(s))
	s =append(s,1,2,3)
	fmt.Println(s)
	// 9 ,12
	//9,10
	//11 , 20
	fmt.Println(len(s),cap(s))
	fmt.Println("===============")
	test()
	fmt.Println("===============")
	test2()
}

func test(){
	s := make([]int,6)
	//s := make([]int,6,10)
	// 6 ,6 ru guo you di san ge  ze cap wei di san ge fou ze dou shi di er ge
	// cap shi zhi jie x2 fan bei
	//6,10
	fmt.Println(len(s),cap(s))
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
	// 9 ,12
	//9,10
	//11 , 20
	fmt.Println(len(s),cap(s))
}

func test2(){
	s := make([]int, 0)
	fmt.Println(len(s),cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(len(s),cap(s))//3 4
	fmt.Println(s)//[1 2 3]

}