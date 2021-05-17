package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url格式 ampq:// 账号：密码@rabbitmq服务器地址：端口号/vhost
const MQURL = "amqp://guest:guest@192.168.147.137:5672/"

type RabbitMQ struct {
	conn *amqp.Connection

	channel *amqp.Channel

	// 队列名称
	QueueName string

	// 交换机
	Exchange string

	// key
	Key string

	// 链接信息
	MqUrl string
}

// 创建 MQ 实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: MQURL}
	var err error

	// 创建 rabbitmq
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		errString := fmt.Sprintf("%s:%s", message, err)
		log.Fatalf(errString)
		panic(errString)
	}
}

// simple模式
// simple模式step：1.创建simple模式RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// simple模式step：2.simple模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		// 是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 发送信息到队列中
	_ = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentEncoding: "text/plain",
			Body:            []byte(message),
		},
	)
}
