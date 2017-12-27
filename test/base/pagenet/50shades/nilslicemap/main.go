package main

import "fmt"

func main(){
	var s []int
	s = append(s,1)
	fmt.Println(s)

	//panic: assignment to entry in nil map
	//bu shi bian yi yi chang  zai yun xing shi yichang ke debug
	//var m map[string]string
	//m["a"] = "a"
	//fmt.Println(m)


	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	//当dst和src的长度不一样时，只会拷贝较短的部分
	sli := []int{1,2,3,4,5}
	t := []int{11,21,31,41,5,6,7,8,9}
	copy(sli, t) // sli=[11,21,31,41,5]
	// copy(t, sli) // t=[1,2,3,4,5,6,7,8,9]

	//看起来Go好像支持多维的Array和Slice，但不是这样的。尽管可以创建数组的数组或者切片的切片。对于依赖于动态多维数组的数值计算应用而言，Go在性能和复杂度上还相距甚远。
	//你可以使用纯一维数组、“独立”切片的切片，“共享数据”切片的切片来构建动态的多维数组
	a := [][]int{[]int{1,2},[]int{3,4}}
	fmt.Println(a)
}
