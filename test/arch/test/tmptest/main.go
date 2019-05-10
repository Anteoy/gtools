package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	a := "yanz01@jumei.com"
	if strings.HasSuffix(a, "@jumei.com") {
		b := strings.TrimSuffix(a, "@jumei.com")
		fmt.Println(b)
	}

	reg := []string{"a", "b", "c"}
	fmt.Println(strings.Join(reg, ";"))
	timeStr := time.Now().Format("20060102150405")
	fmt.Println(timeStr)
}
