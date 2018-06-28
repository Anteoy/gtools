package main

import "fmt"

func main() {
	s := fmt.Sprintf(`Hi teacher，please check your headset and microphone and get ready for the class!
Course details\n：
Course name：%s
Teaching Object：%s
Courseware：%s`, "aa", "bb", "cc")
	fmt.Println(s)
	ss := "a\nbc"
	fmt.Println(ss)
	// '' only char
	sss := '\n'
	fmt.Println(sss)
}
