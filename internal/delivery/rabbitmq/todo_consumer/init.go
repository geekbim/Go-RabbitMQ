package todo_consumer

import (
	"ubersnap/pkg/logger"

	"github.com/streadway/amqp"
)

const (
	exchangeKind       = "direct"
	exchangeDurable    = true
	exchangeAutoDelete = false
	exchangeInternal   = false
	exchangeNoWait     = false

	queueDurable    = true
	queueAutoDelete = false
	queueExclusive  = false
	queueNoWait     = false

	prefetchCount  = 1
	prefetchSize   = 0
	prefetchGlobal = false

	consumeAutoAck   = false
	consumeExclusive = false
	consumeNoLocal   = false
	consumeNoWait    = false
)

type TodoConsumer struct {
	amqpConn *amqp.Connection
	logger   logger.Logger
}

func NewTodoConsumer(amqpConn *amqp.Connection, logger logger.Logger) *TodoConsumer {
	return &TodoConsumer{amqpConn: amqpConn, logger: logger}
}
