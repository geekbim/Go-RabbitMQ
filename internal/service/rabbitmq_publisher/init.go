package rabbitmq_publisher

import (
	"ubersnap/domain/service"
	"ubersnap/pkg/logger"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const (
	publishMandatory = false
	publishImmediate = false
)

type rabbitMQService struct {
	amqpChan *amqp.Channel
	l        logger.Logger
}

func NewRabbitMQService(conn *amqp.Connection, l logger.Logger, newrelic *newrelic.Application) (service.RabbitMQService, error) {
	amqpChan, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "p.amqpConn.Channel")
	}

	rService := rabbitMQService{
		amqpChan: amqpChan,
		l:        l,
	}

	return &rService, nil
}
