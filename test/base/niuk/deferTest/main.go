package main

import "fmt"
//不会到不符合条件的else里面
func main(){
	if (true) {
		defer fmt.Println("11111111111")
	}else {
		defer fmt.Println("22222222222")
	}
	fmt.Println("333333333333333")
	test(4)
}

func test(a int){
	if a>3 {
		defer fmt.Println("aaaaaa")
	}else{
		defer fmt.Println("bbbbbbb")
	}
}
