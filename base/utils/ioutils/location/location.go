package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	local1, err1 := time.LoadLocation("") //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") //服务器上设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}
	a := GetISO8601TimeStamp(time.Now().UTC())
	aa := GetISO8601TimeStamp(time.Unix(time.Now().Unix(), 0))
	aaa := GetISO8601TimeStamp(time.Unix(1502418006, 0))
	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
	fmt.Println(a)
	fmt.Println(aa)
	fmt.Println(aaa)

}

func GetISO8601TimeStamp(ts time.Time) string {
	t := ts.UTC()
	cc := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Printf("#######################: %s\n", cc)
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
