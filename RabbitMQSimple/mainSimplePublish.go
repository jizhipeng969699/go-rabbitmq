package main

import (
	"fmt"
	"go-rabbitmq/RabbitMQ"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("testSimple")
	rabbitmq.PublishSimple("Hello test!")
	fmt.Println("发送成功!")
}
