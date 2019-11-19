package main

import "go-rabbitmq/RabbitMQ"

func main() {
	testTwo := RabbitMQ.NewRabbitMQRouting("exTest","test_two")
	testTwo.ReceiveRouting()
}
