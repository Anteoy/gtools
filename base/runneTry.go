package main

import (
	"fmt"
	"strings"
)

func main(){
	str :="laoYu老虞"
	for  i:=0;i<len(str);i++ {
		fmt.Println(str[i])
	}
	for  i,s :=  range str {
		fmt.Println(i,"Unicode(",s,") string=",string(s))
	}
	r := []rune(str)
	fmt.Println("rune=",r)
	for i:=0;i<len(r) ; i++ {
		fmt.Println("r[",i,"]=",r[i],"string=",string(r[i]))
	}

	str1 :="laoYuStudyGotrue是否包含某字符串"
	fmt.Println(strings.Contains(str1,"go"))         //false
	fmt.Println(strings.Contains(str1,"Go"))         //true
	fmt.Println(strings.Contains(str1,"laoyu"))      //false
	fmt.Println(strings.Contains(str1,"是"))         //true
	fmt.Println(strings.Contains(str1,""))           //true

	fmt.Println(index("soiejfnsoiejhf","iej",false))
}


//字符串subst在字符串s中的索引位置,ignoreCase表示是否忽略大小写
func index(s string, subst string, ignoreCase bool) int {
	n := len(subst)
	if n == 0 {
		return 0
	}
	//to Lower
	if ignoreCase == true {
		s = strings.ToLower(s)
		subst = strings.ToLower(subst)
	}
	c := subst[0]
	//只有一个字符
	if n == 1 {
		// special case worth making fast
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				return i
			}
		}
		return -1
	}
	//字符数 n > 1
	for i := 0; i+n <= len(s); i++ {
		if s[i] == c && s[i:i+n] == subst {
			return i
		}
	}
	return -1
}