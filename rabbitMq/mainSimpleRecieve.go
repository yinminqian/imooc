package main

import "muke/RabbitMQ"

//消费处理消息
func main() {
	rabbit := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbit.ConsumeSimple()
}
