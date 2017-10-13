package main

import (
	"client"
	"fmt"
	"sync"
)

// ==
type ReceiverA struct{}

func (r *ReceiverA) QueueName() string {
	return "receivera_name"
}
func (r *ReceiverA) RouterKey() string {
	return "test-key"
}
func (r *ReceiverA) OnError(e error) {
	fmt.Printf("ReceiverA onError %v\n", e)
}
func (r *ReceiverA) OnReceive(b []byte) bool {
	fmt.Printf("ReceiverA OnReceive %s\n", string(b))
	return true
}

// 启动并开始处理数据
func main() {
	wg := sync.WaitGroup{}
	go a(wg)
	go a(wg)
	go a(wg)
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)
	wg.Wait()
}

func a(wg sync.WaitGroup) {
	defer func() { wg.Add(-1) }()
	// 假设这里有一个AReceiver和BReceiver
	aReceiver := &ReceiverA{}

	mq := irabbit.NewRabbitConsumer("", "test-exchange", "direct", 3)
	// 将这个接收者注册到
	mq.RegisterReceiver(aReceiver)
	mq.Start()
}
