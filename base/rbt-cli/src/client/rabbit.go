package irabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

// Receiver 接口
type Receiver interface {
	QueueName() string // 接收者监听队列
	RouterKey() string // 队列绑定的路由
	OnError(error)     // 错误处理
	OnReceive([]byte) bool
}

// RabbitMQ 用于管理和维护rabbit mq连接
type RabbitMQ struct {
	ServerAddr string
	MaxRefresh int
	channel    *amqp.Channel
}

// RabbitConsumerManager consumer 管理
type RabbitConsumerManager struct {
	RabbitMQ
	wg            sync.WaitGroup
	receivers     []Receiver
	exchangeName  string // exchange的名称
	exchangeType  string // exchange的类型
	ConsumerMount int
}

// RabbitProducerManager 消息生产者管理
type RabbitProducerManager struct {
	RabbitMQ
	exchangeName string // exchange的名称
	exchangeType string // exchange的类型
	RoutingKey   string
	Reliable     bool
}

// Publish 发送mq 消息
func (p *RabbitProducerManager) Publish(msg string) error {
	i := 1
	for !p.Refresh() {
		if i > p.MaxRefresh {
			return fmt.Errorf("mq msg 无法发送: 超过最大重连次数%d仍无法连接,请人工排查错误\n", p.MaxRefresh)
		}
		log.Printf("producer 获取rabbit连接失败,将在5s后进行第%d次重试,自定义最大重试次数为%d次\n", i, p.MaxRefresh)
		time.Sleep(5 * time.Second)
		i = i + 1
	}
	if err := p.channel.ExchangeDeclare(
		p.exchangeName,
		p.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("Exchange Declare 定义exchange出错了: %s\n", err)
	}
	//保证消息可靠
	if p.Reliable {
		log.Printf("开启发布确认.")
		if err := p.channel.Confirm(false); err != nil {
			return fmt.Errorf("Channel 无法切换到确认模式: %s\n", err)
		}

		confirms := p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer confirmOne(confirms)
	}
	log.Printf("exchange定义完成, 发布 %dB 消息 (%q)", len(msg), msg)
	if err := p.channel.Publish(
		p.exchangeName,
		p.RoutingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(msg),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("exchange 发布出错: %s\n", err)
	}
	return nil
}

// confirmOne 发布确认
func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Printf("等待发布确认")

	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("发布成功: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("发布失败: %d", confirmed.DeliveryTag)
	}
}

// New 创建一个新的操作Rabbit Producer的对象
func NewRabbitProducer(ServerAddr, exchangeName, exchangeType, routingKey string, reliable bool, retry int) *RabbitProducerManager {
	// 这里可以根据自己的需要去定义
	instanse := &RabbitProducerManager{
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		Reliable:     reliable,
		RoutingKey:   routingKey,
	}
	instanse.ServerAddr = ServerAddr
	instanse.MaxRefresh = retry
	return instanse
}

// New 创建一个新的操作Rabbit Consumer的对象
func NewRabbitConsumer(ServerAddr, exchangeName, exchangeType string, consumerMount int) *RabbitConsumerManager {
	// 这里可以根据自己的需要去定义
	instanse := &RabbitConsumerManager{
		exchangeName: exchangeName,
		exchangeType: exchangeType,
	}
	instanse.ServerAddr = ServerAddr
	instanse.ConsumerMount = consumerMount
	return instanse
}

// prepareExchange
func (mq *RabbitConsumerManager) prepareExchange() error {
	err := mq.channel.ExchangeDeclare(
		mq.exchangeName,
		mq.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		return err
	}
	return nil
}

// run 开始获取连接并初始化相关操作
func (mq *RabbitConsumerManager) run() {
	if !mq.Refresh() {
		fmt.Errorf("rabbit刷新连接失败，即将进行重连: %s", mq.ServerAddr)
		return
	}

	fmt.Println("rabbitMQ 连接成功,开启消费消息服务...")
	// 初始化Exchange
	mq.prepareExchange()

	for _, receiver := range mq.receivers {
		mq.wg.Add(1)
		fmt.Printf("消费者消费队列 :%s,开始进入监听状态...\n", receiver.QueueName())
		for i := 0; i < mq.ConsumerMount; i++ {
			go mq.listen(receiver)
		}
	}

	mq.wg.Wait()

	log.Println("所有处理queue的任务都意外退出了")

	//关闭当前连接
	mq.Distory()
}

func (mq *RabbitConsumerManager) Distory() {
	mq.channel = nil
}

// Refresh 初始和重试连接
func (mq *RabbitMQ) Refresh() bool {
	rabbitmqConn, err := amqp.Dial(mq.ServerAddr)
	if err != nil {
		panic("RabbitMQ 初始化失败: " + err.Error())
		return false
	}
	mq.channel, err = rabbitmqConn.Channel()
	if err != nil {
		panic("打开Channel失败: " + err.Error())
		return false
	}
	return true
}

// Start 启动Rabbitmq的客户端
func (mq *RabbitConsumerManager) Start() {
	i := 0
	for {
		mq.run()
		// 消费端重连机制
		log.Printf("获取rabbit连接失败,将在5s后进行第%d次重试\n", i)
		time.Sleep(5 * time.Second)
		i = i + 1
	}
}

// RegisterReceiver 加入注册的消费者
func (mq *RabbitConsumerManager) RegisterReceiver(receiver Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

// Listen 监听指定路由发来的消息
func (mq *RabbitConsumerManager) listen(receiver Receiver) {
	defer mq.wg.Done()

	// 获取每个接收者需要监听的队列和路由
	queueName := receiver.QueueName()
	routerKey := receiver.RouterKey()

	// 声明Queue
	_, err := mq.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		// 当队列初始化失败的时候，需要告诉这个接收者相应的错误
		receiver.OnError(fmt.Errorf("初始化队列 %s 失败: %s", queueName, err.Error()))
	}

	// 将Queue绑定到Exchange上去
	err = mq.channel.QueueBind(
		queueName,
		routerKey,
		mq.exchangeName,
		false,
		nil,
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("绑定队列 [%s - %s] 到交换机失败: %s", queueName, routerKey, err.Error()))
	}

	// 获取消费通道 确保rabbitmq会一个一个发消息
	mq.channel.Qos(1, 0, true)
	deliveries, err := mq.channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error()))
	}

	for msg := range deliveries {
		for !receiver.OnReceive(msg.Body) {
			log.Println("receiver 数据处理失败，即将进行重试操作")
			time.Sleep(1 * time.Second)
		}
		// 确认收到本条消息, multiple必须为false
		msg.Ack(false)
	}
}
