package rabbitmq_publisher

import (
	"fmt"
	"time"
	"ubersnap/pkg/service/rabbitmq"
	"ubersnap/pkg/utils"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func (p *rabbitMQService) publishToQueue(body []byte, exchange rabbitmq.RabbitExchangeType, queueName rabbitmq.RabbitQueueType) error {
	p.l.Info(fmt.Sprintf("Publishing message Exchange: %s, RoutingKey: %s", exchange, queueName))
	p.l.Info(string(body))

	if err := p.amqpChan.Publish(
		string(exchange),
		string(queueName),
		publishMandatory,
		publishImmediate,
		amqp.Publishing{
			Type:         string(exchange),
			ContentType:  string(rabbitmq.RABBIT_CONTENT_TYPE_JSON),
			DeliveryMode: amqp.Persistent,
			MessageId:    utils.GenerateUUID(),
			Timestamp:    time.Now(),
			Body:         body,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}

	return nil
}
