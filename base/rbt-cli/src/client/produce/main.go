package main

import (
	"client"
	"log"
	"strconv"
)

func main() {
	producer := irabbit.NewRabbitProducer("amqp://", "test-exchange", "direct", "test-key", true, 5)
	for i := 0; i < 10000; i++ {
		msg := "foobar ribenren test" + strconv.Itoa(i)
		if err := producer.Publish(msg); err != nil {
			log.Fatalf("%s", err)
		}
		log.Printf("published %dB OK", len(msg))
	}
}
