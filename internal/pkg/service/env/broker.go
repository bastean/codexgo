package env

import (
	"os"
)

var (
	BrokerRabbitMQURI  = os.Getenv("CODEXGO_BROKER_RABBITMQ_URI")
	BrokerRabbitMQName = os.Getenv("CODEXGO_BROKER_RABBITMQ_NAME")
)
