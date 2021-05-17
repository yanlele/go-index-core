package RabbitMQ

import "github.com/streadway/amqp"

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
func NewRabbitMQ(quenuName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: quenuName, Exchange: exchange, Key: key, MqUrl: MQURL}
	var err error

	// 创建 rabbitmq
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err,"创建连接错误")
	rabbitmq.channel,err= rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err,"获取channel失败")
	return rabbitmq
}


