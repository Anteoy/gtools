package main

import (
	"fmt"
)
//大端序的存储方式是高位字节存储在低地址上；小端序的存储方式是高位字节存储在高地址上。本人的机器是按小端序来存储的，所以gid在我的内存上的存储序列是这样的：0x78, 0x56, 0x34, 0x12。如果您的机器是按大端序来存储，则gid的存储序列刚好反过来：0x12, 0x34, 0x56, 0x78。对于强制转换后的uid，肯定是产生了截断行为。因为uid只占1个字节，转换后的结果必然会丢弃掉多余的3个字节。截断的规则是：保留低地址上的数据，丢弃多余的高地址上的数据
func IsBigEndian() bool {
	var i int32 = 0x12345678
	var b byte  = byte(i) //使用int32也可以
	if b == 0x12 {
		return true
	}

	return false
}

func main() {
	if IsBigEndian() {
		fmt.Println("大端序")
	} else {
		fmt.Println("小端序")
	}
}