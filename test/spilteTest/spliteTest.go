package main

import (
	"strings"
	"fmt"
)

func main()  {
	var a string = "12班"
	var aa string = "1班"
	// len2 cap2 ""
	b := strings.Split(a,"班")
	bb := strings.Split(aa,"班")
	fmt.Printf("%+v\n",b)
	fmt.Printf("%+v\n",bb)
}