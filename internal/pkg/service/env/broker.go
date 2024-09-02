package env

import (
	"os"
)

var (
	BrokerRabbitMQURI, BrokerRabbitMQName string
)

func Broker() {
	BrokerRabbitMQURI = os.Getenv("CODEXGO_BROKER_RABBITMQ_URI")
	BrokerRabbitMQName = os.Getenv("CODEXGO_BROKER_RABBITMQ_NAME")
}
