package main

import "fmt"

func ByteSlice(b []byte) []byte {
	return b
}
func main() {
	b := []byte{97, 98}
	c := []byte{65, 66}
	var d byte = 66
	var dd byte = 66
	var e uint8 = 49
	// u8 := []uint8{2, 3}
	var u8 []uint8
	u8 = b
	fmt.Printf("%T %T\n", b, u8)
	fmt.Println(ByteSlice(b))
	fmt.Println(ByteSlice(u8))
	fmt.Println(string(b))
	fmt.Println(string(c))
	fmt.Println(string(d))
	fmt.Println(string(dd))
	fmt.Println(string(e))
	fmt.Println("======================")

	test := make([]byte, 2)
	test[0] = 65
	test[1] = 97
	fmt.Println(string(test[0]))
	fmt.Println(string(test))
	fmt.Println(string(test[1]))
}
