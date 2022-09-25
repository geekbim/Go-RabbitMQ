package todo_consumer

import (
	"context"
	"fmt"
	"ubersnap/pkg/service/rabbitmq"

	"github.com/streadway/amqp"
)

func (c *TodoConsumer) worker(ctx context.Context, messages <-chan amqp.Delivery) {

	for delivery := range messages {

		c.logger.Info(fmt.Sprintf("processDeliveries deliveryTag% v", delivery.DeliveryTag))
		c.logger.Info(string(delivery.Body))

		payloads := rabbitmq.Todo{}
		_ = payloads.UnSerialize(delivery.Body)

		// put usecase here
	}

	c.logger.Info("Deliveries channel closed")
}
