package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main()  {
	var a string = "12班"
	var aa string = "1班"
	// len2 cap2 ""
	b := strings.Split(a,"班")
	c ,_:= strconv.ParseInt(b[0],10,64)
	bb := strings.Split(aa,"班")
	fmt.Printf("%+v\n",b)
	fmt.Printf("%+v\n",bb)
	d:= "dsj"+string(3)+string(4)
	e := string(3)
	f := fmt.Sprintf("%sg%dc%d","dkjshf",4,3)
	fmt.Printf("%+v\n",e)
	fmt.Printf("%+v\n",c)
	fmt.Printf("%+v\n",d)
	fmt.Printf("%+v\n",f)

}