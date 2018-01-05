package main

// not string to bytes ,for pointer
func main() {
	s := "123"
	ps := &s //123
	b := []byte(*ps) //123
	pb := &b //123

	s += "4" //1234
	*ps += "5" //1235
	b[1] = '0' //103

	println(*ps) //1235 (xx) => 12345
	println(string(*pb)) //103
}