package communication

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service"
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	exchange string
}

func (rmq *RabbitMQ) AddExchange(exchange *exchange.Exchange) {
	err := rmq.Channel.ExchangeDeclare(
		exchange.Name,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	rmq.exchange = exchange.Name

	service.FailOnError(err, "failed to declare an exchange")
}

func (rmq *RabbitMQ) AddQueue(queue *queue.Queue) {
	_, err := rmq.Channel.QueueDeclare(
		queue.Name,
		true,
		false,
		false,
		false,
		nil,
	)

	service.FailOnError(err, "failed to declare a queue")
}

func (rmq *RabbitMQ) AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) {
	for _, bindingKey := range bindingKeys {
		log.Printf("binding queue %s to exchange %s with routing key %s",
			queue.Name, rmq.exchange, bindingKey)

		err := rmq.Channel.QueueBind(
			queue.Name,
			bindingKey,
			rmq.exchange,
			false,
			nil)

		service.FailOnError(err, "failed to bind a queue")
	}
}

func (rmq *RabbitMQ) AddQueueConsumer(consumer model.Consumer) {
	for _, queue := range consumer.SubscribedTo() {
		messages, err := rmq.Channel.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)

		service.FailOnError(err, "failed to register a consumer")

		go func() {
			for delivery := range messages {
				message := new(message.Message)

				err := json.Unmarshal(delivery.Body, message)

				service.FailOnError(err, "failed to delivery a message")

				consumer.On(message)

				delivery.Ack(false)
			}
		}()
	}
}

func (rmq *RabbitMQ) PublishMessages(messages []*message.Message) {
	for _, message := range messages {
		if message.Id == "" {
			message.Id = uuid.NewString()
		}

		if message.OccurredOn == "" {
			message.OccurredOn = time.Now().UTC().String()
		}

		messageJson, _ := json.Marshal(message)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := rmq.Channel.PublishWithContext(ctx,
			rmq.exchange,
			message.Type,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         messageJson,
			})

		service.FailOnError(err, "failed to publish a event")
	}
}

func CloseRabbitMQ(rmq *RabbitMQ) {
	err := rmq.Channel.Close()

	if err != nil {
		service.FailOnError(err, "failed to close channel")
	}

	err = rmq.Connection.Close()

	if err != nil {
		service.FailOnError(err, "failed to close rabbitmq connection")
	}
}

func NewRabbitMQ(uri string) model.Broker {
	conn, err := amqp.Dial(uri)
	service.FailOnError(err, "failed to connect to rabbitmq")

	ch, err := conn.Channel()
	service.FailOnError(err, "failed to open a channel")

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
	}
}
