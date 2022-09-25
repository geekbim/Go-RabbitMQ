package rabbitmq

import (
	"log"
	"ubersnap/pkg/config"

	"github.com/streadway/amqp"
)

func NewRabbitMQConn() (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.RabbitMQ().Url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
