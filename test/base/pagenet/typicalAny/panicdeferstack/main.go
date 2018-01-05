package main

func f1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}

func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}

func f3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

// f1-begin f2-begin f3-begin f3-end f2-end f1-end  panic (try no no no)
func main() {
	f1()
}
