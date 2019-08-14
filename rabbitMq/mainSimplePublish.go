package main

import (
	"fmt"
	"muke/RabbitMQ"
)

//生产消息
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.PublishSimple("这是一条消息")
	fmt.Print("消息发送完毕")
}
