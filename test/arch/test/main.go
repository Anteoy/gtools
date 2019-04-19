package main

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {

		select {
		case n := <-data:
			println(n)
		}

		for v := range data {
			println(v)
		}

		println("receive over")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	data <- 4
	// data <- 3
	//close(data)

	println("send over")
	<-exit
}
