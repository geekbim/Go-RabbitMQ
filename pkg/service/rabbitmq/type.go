package rabbitmq

type (
	RabbitQueueType    string
	RabbitExchangeType string
	ConsumerTagType    string
)

const (
	RABBIT_Q_CREATE_TODO  RabbitQueueType = "req.create.todo"
	RABBIT_Q_TODO_CREATED RabbitQueueType = "todo.created"
	RABBIT_Q_DELETE_TODO  RabbitQueueType = "req.delete.todo"
	RABBIT_Q_TODO_DELETED RabbitQueueType = "todo.deleted"

	RABBIT_E_CREATE_TODO  RabbitExchangeType = "req.create.todo"
	RABBIT_E_TODO_CREATED RabbitExchangeType = "todo.created"
	RABBIT_E_DELETE_TODO  RabbitExchangeType = "req.delete.todo"
	RABBIT_E_TODO_DELETED RabbitExchangeType = "todo.deleted"

	RABBIT_CT_TODO ConsumerTagType = "todo_tag"
)
