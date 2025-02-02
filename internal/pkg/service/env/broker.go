package env

import (
	"os"
)

var (
	BrokerRabbitMQURI, BrokerRabbitMQName string
)

func Broker() {
	BrokerRabbitMQURI = os.Getenv(BROKER_RABBITMQ_URI)
	BrokerRabbitMQName = os.Getenv(BROKER_RABBITMQ_NAME)
}

func HasRabbitMQ() bool {
	return BrokerRabbitMQURI != ""
}
