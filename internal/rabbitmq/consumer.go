package rabbitmq

import (
	"errors"

	ampq "github.com/rabbitmq/amqp091-go"
)

type (
	Consumer struct {
		connection *ampq.Connection
	}
	ConsumerOptions struct {
		ExchangeName string
		ExchangeType string
		QueueName    string
		Callback     func(msg any) error
	}
)

func NewConsumer(connection *ampq.Connection) (Consumer, error) {
	return Consumer{connection}, nil
}

func (c Consumer) Consume(ops ConsumerOptions) error {
	if err := c.validate(ops); err != nil {
		return err
	}
	channel, err := c.connection.Channel()
	if err != nil {
		return err
	}
	err = channel.ExchangeDeclare(ops.ExchangeName, ops.ExchangeType, false, false, false, false, nil)
	if err != nil {
		return err
	}
	queue, err := channel.QueueDeclare(ops.QueueName, false, false, false, false, nil)
	if err != nil {
		return err
	}
	err = channel.QueueBind(ops.QueueName, "", ops.ExchangeName, false, nil)
	if err != nil {
		return err
	}
	msgs, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	//todo close the channel
	go func() {
		for msg := range msgs {
			if err := ops.Callback(msg); err != nil {
				return
			}
			if err := msg.Ack(false); err != nil {
				return
			}
		}
	}()
	return nil
}

func (c Consumer) validate(ops ConsumerOptions) error {
	if ops.ExchangeType == "" {
		return errors.New("exchange type is required")
	}
	if ops.ExchangeName == "" {
		return errors.New("exchange name is required")
	}
	if ops.QueueName == "" {
		return errors.New("queue name is required")
	}
	if ops.Callback == nil {
		return errors.New("callback is required")
	}
	return nil
}
