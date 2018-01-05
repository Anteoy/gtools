package main


import "fmt"

func get() []byte {
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
	return raw[:3]
}
func main() {
	data := get()
	// zhan zhuo mao keng bu hui bei la ji hui shou
	fmt.Println(len(data),cap(data),&data[0]) //prints: 3 10000 <byte_addr_x>

	fmt.Println("======")
	//为了避免这个陷阱，你需要从临时的slice中拷贝数据（而不是重新划分slice）。
	data = get2()
	fmt.Println(len(data),cap(data),&data[0]) //prints: 3 3 <byte_addr_y>
}

func get2() []byte {
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
	res := make([]byte,3)
	copy(res,raw[:3])
	return res
}
