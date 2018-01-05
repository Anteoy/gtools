package main

import "fmt"

func main() {
	data := []int{1,2,3}
	for _,v := range data {
		v *= 10 //original item is not changed
		v = v*10
	}
	fmt.Println("data:",data) //prints data: [1 2 3]


	//如果你需要更新原有集合中的数据，使用索引操作符来获得数据。
	for i,_ := range data {
		data[i] *= 10
	}
	fmt.Println("data:",data) //prints data: [10 20 30]

	//如果你的集合保存的是指针，那规则会稍有不同。
	//如果要更新原有记录指向的数据，你依然需要使用索引操作，但你可以使用for range语句中的第二个值来更新存储在目标位置的数据。
	type aa struct{
		num int
	}
	data1 := []*aa { {1},{2},{3} }
	for _,v := range data1 {
		v.num *= 10
	}
	fmt.Println(data1[0],data1[1],data1[2]) //prints &{10} &{20} &{30}
}
