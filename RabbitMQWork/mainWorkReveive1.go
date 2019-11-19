package main

import "go-rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("testSimple")
	rabbitmq.ConsumeSimple()
}
