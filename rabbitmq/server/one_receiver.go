package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 处理错误
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//var msgCh = make(chan string)
//var fovever = make(chan bool)

func mqhander() {

	conn, err := amqp.Dial("amqp://user:123456@10.158.215.21:5672/")
	FailOnError(err, "连接 RabbitMQ 失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"botkube",
		true,
		false,
		false,
		false,
		nil,
	)
	err = ch.Qos(1, 0, false)
	FailOnError(err, "声明队列失败")

	msgs, err := ch.Consume(q.Name,
		"", true, false, false, false, nil)
	FailOnError(err, "注册消费者失败")

	//fovever := make(chan bool)
	var msg string
	//go func() {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		msg = string(d.Body)
		//msgCh <- msg
	}
	//}()

	fmt.Println(msg)
}
func main() {
	mqhander()
	fmt.Println("ssssssssssss")
	//<-fovever
	//fmt.Println(len(msgCh))
	//for range msgCh {
	//	a := <-msgCh
	//	fmt.Println(a)
	//	fmt.Println("hi")
	//}
}
