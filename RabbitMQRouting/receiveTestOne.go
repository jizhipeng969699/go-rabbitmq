package main

import "go-rabbitmq/RabbitMQ"

func main() {
	testOne := RabbitMQ.NewRabbitMQRouting("exTest","test_one")
	testOne.ReceiveRouting()
}
