package main

import (
	"fmt"
	"go-rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	testOne := RabbitMQ.NewRabbitMQRouting("exTest", "test_one")
	testTwo := RabbitMQ.NewRabbitMQRouting("exTest", "test_two")
	for i := 0; i <= 10; i++ {
		testOne.PublishRouting("Hello test one:" + strconv.Itoa(i))
		testTwo.PublishRouting("Hello test two:" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println("publish:",i)
	}
}
