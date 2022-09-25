package config

import (
	"os"
)

type RabbitMQConfig struct {
	Url string
}

func RabbitMQ() *RabbitMQConfig {
	return &RabbitMQConfig{
		Url: os.Getenv("RBMQ_URL"),
	}
}
