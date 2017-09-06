package main

import (
	"fmt"
	"container/list"
	"strconv"
)

func main() {
	s := []int{11, 22, 33, 44, 55, 66}
	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(strconv.Itoa(i))
	}
	if contain, e := Contains(l, "1"); contain {
		l.Remove(e)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

//不过从slice这种数据结构来看，本身并不适合做删除操作。抛开语言，只谈数据结构，我们知道数组删除是会移动元素的，效率会比较低。当然任何语言的数组实现（顺序存储），删除元素都避免不了移动元素。
//所以，如果会平凡删除中间或开头的元素，更好的是选择链表这样的数据结构。Go中可以使用map或container/list包。
func Contains(l *list.List, value string) (bool, *list.Element) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true, e
		}
	}
	return false, nil
}