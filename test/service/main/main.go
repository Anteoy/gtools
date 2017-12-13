package main

import (
	"fmt"
	"math"
)

func main(){
	a := []int{1,12,21,31}
	b := []int{}
	b = a
	for i:=0;i< len(a)-1;i++{
		fmt.Println(a[i])
		fmt.Println(a[i+1])
		//if a[i+1] - a[i] >= 7{
			//if 7 <= (a[i+1]-a[i]) && (a[i+1]-a[i]) <= 10{
			//	b = append(b,a[i]+5)
			//}else {
			//	//add
			//	count := int((a[i+1] - a[i] - 5)/5)
			//	for j:=0;j<count;j++ {
			//		b = append(b,a[i]+5*(j+1))
			//	}
			//}
			if float64(a[i+1])/float64(a[i]) > 1{
				countflo :=(float64(a[i+1])-float64(a[i]))/5
				if math.Ceil(countflo) == math.Floor(countflo) {
					countflo -= 1
				}
				count := int(countflo)
				for j:=0;j<count;j++ {
					b = append(b,a[i]+5*(j+1))
				}
			}
		//}
	}
	fmt.Println("====")
	for _,b :=range b {
		fmt.Println(b)
	}

	//fmt.Println("test")
	//c := float32(12)/float32(5)
	//fmt.Println(c)

	//aa := math.Ceil(5.0)
	//bb := math.Floor(5.0)
	//fmt.Println(aa)
	//fmt.Println(bb)
}
