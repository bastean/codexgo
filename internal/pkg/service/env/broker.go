package env

import (
	"os"
)

var (
	BrokerRabbitMQURI  = os.Getenv("BROKER_RABBITMQ_URI")
	BrokerRabbitMQName = os.Getenv("BROKER_RABBITMQ_NAME")
)
