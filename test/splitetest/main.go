package main

import (
	"strings"
	"fmt"
)

func main(){
	msg := "201802081339330DYtzbwKUNZ7cetl94^===&d1^===&3"
	msgs := strings.Split(msg,"^===&")
	fmt.Printf("%+v",msgs)
}
