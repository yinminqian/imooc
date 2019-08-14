package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式 amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imoocuser"

type RabbitMQ struct {
	coon    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//链接名称
	Mqurl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queuename string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queuename,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	//创建 rabbitmq链接
	rabbitmq.coon, err = amqp.Dial(rabbitmq.Mqurl)
	//处理错误
	rabbitmq.failOnErr(err, "创建链接错误!")
	rabbitmq.channel, err = rabbitmq.coon.Channel()
	rabbitmq.failOnErr(err, "获取channel失败!")
	return rabbitmq
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.coon.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatal("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//简单模式Step1: 简单模式下的rabbitMQ实例
func NewRabbitMQSimple(queuename string) *RabbitMQ {
	return NewRabbitMQ(queuename, "", "")
}

//简单模式Step2: 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列,如果丢列不存在会自动创建,如果存在则跳过创建
	//保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否排他性 (仅自己可见)
		false,
		//是否阻塞
		false,
		//其他属性
		nil,
	)

	if err != nil {
		fmt.Print(err)
	}
	//发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true,根据exhange和routkey规则,如果无法找到符合条件的对列,那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange发送消息到对列后发现对列上没有绑定消费者,则会把消息发还给发送者
		false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message),})
}

//接受处理消息 消费消息
//生产消息和消费消息都要先
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列,如果丢列不存在会自动创建,如果存在则跳过创建
	//保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否排他性 (仅自己可见)
		false,
		//是否阻塞
		false,
		//其他属性
		nil,
	)

	if err != nil {
		fmt.Print(err)
	}

	msgs, err := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true,表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//队列消费是否阻塞
		false, nil,
	)
	if err != nil {
		fmt.Print(err)
	}
	forver := make(chan bool)
	//启用协程处理消息

	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("Received a meaasge: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages,To exit press CTRL + C ")
	<-forver
}
