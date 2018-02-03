package main

import (
	"fmt"
	"strconv"
	"encoding/hex"
	"crypto/md5"
)

func main() {
	r1 := HexStr([]byte(string(176028871861)))
	r2 := HexStr([]byte(string(176028871862)))
	s1 := strconv.FormatInt(176028871861, 10)
	s2 := strconv.FormatInt(176028871862, 10)
	r3 := HexStr([]byte(s1))
	r4 := HexStr([]byte(s2))
	fmt.Printf("%s\n%s", r1, r2)
	fmt.Println("======")
	fmt.Printf("%s\n%s", r3, r4)

}

func HexStr(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
