package service

import (
	"ubersnap/pkg/service/rabbitmq"
)

type RabbitMQService interface {
	PublishTodos(payload rabbitmq.Todo) error
}
