package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4

// 1. xian zhi xing han shu : calc() cong shang dao xia
// 2. defer cong xia dao shang
// 3. bian liang de fan  wei zai dai ma shang mian ,ying gai shi fuben kao bei le han shu de zhi ,suo yi buhui shou xia mian ying xiang
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
