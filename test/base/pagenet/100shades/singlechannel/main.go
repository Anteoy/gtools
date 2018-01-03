package main

//dan xiang channel shi wei le yue su
func main() {
	c := make(chan int, 3)

	var send chan<- int = c
	var recv <-chan int = c

	send <- 1

	println(<-recv)

}
