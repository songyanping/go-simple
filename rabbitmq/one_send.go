package main

import (
	"github.com/streadway/amqp"
	"log"
)

// 处理错误
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://user:123456@10.158.215.21:5672/")
	FailOnError(err, "连接 RabbitMQ 失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	// 声明队列，存储消息
	q, err := ch.QueueDeclare(
		"botkube",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "声明队列失败")

	body := "mq msg4444444444"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf("[x] Sent %s", body)
	FailOnError(err, "发送消息失败")
}
