package main

/*
 enum Color { BLACK = 10, RED, BLUE };
 typedef enum { INSERT = 3, DELETE } Mode;
*/
import "C"
import "fmt"

func main() {
	var c C.enum_Color = C.RED
	var x uint32 = c
	fmt.Println(c, x)
	var m C.Mode = C.INSERT
	fmt.Println(m)

	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", m)
}
