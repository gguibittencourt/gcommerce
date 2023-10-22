package rabbitmq

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func NewConnection() (*ampq.Connection, error) {
	return ampq.Dial("amqp://guest:guest@localhost:5672/")
}
