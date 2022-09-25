package rabbitmq_publisher

import (
	"ubersnap/pkg/service/rabbitmq"
)

func (p *rabbitMQService) PublishTodo(payload rabbitmq.Todo) error {
	body, _ := payload.Serialize()
	return p.publishToQueue(body, rabbitmq.RABBIT_E_CREATE_TODO, rabbitmq.RABBIT_Q_CREATE_TODO)
}
