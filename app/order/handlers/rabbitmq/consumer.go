package rabbitmq

import (
	"fmt"

	ampq "github.com/rabbitmq/amqp091-go"
)

type (
	Service interface {
		Consume(out chan ampq.Delivery) error
	}
)

func NewConsumer(service Service) error {
	go func() {
		msgs := make(chan ampq.Delivery)
		err := service.Consume(msgs)
		for msg := range msgs {
			fmt.Printf("msg consumed: %v \n", msg)
			msg.Ack(false)
		}
		fmt.Println(err)
	}()
	return nil
}
