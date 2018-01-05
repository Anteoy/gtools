package main

import (
	"fmt"
	"bytes"
)
func main() {
	//path := []byte("AAAA/BBBBBBBBB")
	//sepIndex := bytes.IndexByte(path,'/')
	//dir1 := path[:sepIndex]
	//dir2 := path[sepIndex+1:]
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	////结果与你想的不一样。与"AAAAsuffix/BBBBBBBBB"相反，你将会得到"AAAAsuffix/uffixBBBB"。
	////这个情况的发生是因为两个文件夹的slice都潜在的引用了同一个原始的路径slice。这意味着原始路径也被修改了。
	//dir1 = append(dir1,"suffix"...)
	//path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	//fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	//fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB (not ok)
	//fmt.Println("new path =>",string(path))

	//通过分配新的slice并拷贝需要的数据，你可以修复这个问题。另一个选择是使用完整的slice表达式。
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex:sepIndex] //full slice expression
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1,"suffix"...)
	path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)
	fmt.Println("new path =>",string(path))
}
