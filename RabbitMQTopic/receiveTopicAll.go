package main

import "go-rabbitmq/RabbitMQ"

func main() {
	testOne := RabbitMQ.NewRabbitMQTopic("exTestTopic","#")
	testOne.ReceiveTopic()
}
