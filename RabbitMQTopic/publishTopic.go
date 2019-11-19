package main

import (
	"fmt"
	"go-rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	testOne := RabbitMQ.NewRabbitMQTopic("exTestTopic", "test.topic.one")
	testTwo := RabbitMQ.NewRabbitMQTopic("exTestTopic", "test.topic.two")
	for i := 0; i <= 10; i++ {
		testOne.PublishTopic("Hello test topic one:" + strconv.Itoa(i))
		testTwo.PublishTopic("Hello test topic two:" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println("publish:", i)
	}
}
