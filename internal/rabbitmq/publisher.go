package rabbitmq

import (
	"context"
	"encoding/json"

	ampq "github.com/rabbitmq/amqp091-go"
)

type (
	Publisher struct {
		connection *ampq.Connection
	}
	PublishOptions struct {
		ExchangeName string
		Msg          any
	}
)

func NewPublisher(connection *ampq.Connection) (Publisher, error) {
	return Publisher{connection}, nil
}

func (p Publisher) Publish(ctx context.Context, exchangeName string, msg any) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.PublishWithContext(
		ctx,
		exchangeName,
		"",
		false,
		false,
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}
