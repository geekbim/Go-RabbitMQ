package config

import (
	"os"
	"strconv"
	"time"
	"ubersnap/pkg/config"
	"ubersnap/pkg/service/rabbitmq"
)

type ServerConfig struct {
	Port     string
	RabbitMQ RabbitMQ
	TimeOut  time.Duration
}

type RabbitMQ struct {
	Url                string
	Exchange           RExchange
	Queue              RQueue
	RoutingKeyEmail    string
	RoutingKeySms      string
	RoutingKeyEligible string
	ConsumerTag        string
	WorkerPoolSize     int
}

type RQueue struct {
	CreateTodo  string
	TodoCreated string
	DeleteTodo  string
	TodoDeleted string
}

type RExchange struct {
	Default            string
	Reward             string
	Survey             string
	EligibleRespondent string
}

func convertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Server() ServerConfig {
	cfg := ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
		RabbitMQ: RabbitMQ{
			Url:      config.RabbitMQ().Url,
			Exchange: RExchange{
				// Reward:             string(rabbitmq.RABBIT_E_REWARD),
				// Default:            string(rbmq.RABBIT_E_BROADCAST),
				// Survey:             string(rbmq.RABBIT_E_SURVEY),
				// EligibleRespondent: string(rbmq.RABBIT_E_ELIGIBLE_RESPONDENT),
			},
			Queue: RQueue{
				// Sms:                string(rbmq.RABBIT_Q_SMS),
				// SmsDirect:          string(rbmq.RABBIT_Q_SMS_DIRECT),
				// Email:              string(rbmq.RABBIT_Q_EMAIL),
				// Reward:             string(rbmq.RABBIT_Q_REWARD),
				// EligibleRespondent: string(rbmq.RABBIT_Q_ELIGIBLE_RESPONDENT_V2),
				// CroxSms:            string(rbmq.RABBIT_Q_CROX_SMS),
				CreateTodo:  string(rabbitmq.RABBIT_Q_CREATE_TODO),
				TodoCreated: string(rabbitmq.RABBIT_Q_TODO_CREATED),
				DeleteTodo:  string(rabbitmq.RABBIT_Q_DELETE_TODO),
				TodoDeleted: string(rabbitmq.RABBIT_Q_TODO_DELETED),
			},
			ConsumerTag:    string(rabbitmq.RABBIT_CT_TODO),
			WorkerPoolSize: 1,
		},
		TimeOut: time.Duration(convertInt("APP_TIMEOUT")) * time.Second,
	}
	err := cfg.Validate()
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *ServerConfig) Validate() error {
	fields := []string{
		"SERVER_PORT",
		"RBMQ_URL",
	}

	for _, f := range fields {
		err := config.Required(f)
		if err != nil {
			return err
		}
	}
	return nil
}
