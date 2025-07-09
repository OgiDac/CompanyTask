package publisher

import (
	"encoding/json"

	"github.com/OgiDac/CompanyTask/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitPublisher struct {
	channel *amqp.Channel
	queue   string
}

func NewRabbitPublisher(channel *amqp.Channel, queue string) domain.EventPublisher {
	return &rabbitPublisher{
		channel: channel,
		queue:   queue,
	}
}

func (r *rabbitPublisher) PublishEvent(envelope domain.UserEventEnvelope) error {
	body, err := json.Marshal(envelope)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		r.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}
