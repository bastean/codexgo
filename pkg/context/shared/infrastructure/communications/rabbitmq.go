package communications

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	models.Logger
	exchange string
}

func (rmq *RabbitMQ) AddRouter(router *messages.Router) error {
	err := rmq.Channel.ExchangeDeclare(
		router.Name,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	rmq.exchange = router.Name

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "AddRouter",
			What:  "failure to declare a router",
			Why: errors.Meta{
				"Router": router.Name,
			},
			Who: err,
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueue(queue *messages.Queue) error {
	_, err := rmq.Channel.QueueDeclare(
		queue.Name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "AddQueue",
			What:  "failure to declare a queue",
			Why: errors.Meta{
				"Queue": queue.Name,
			},
			Who: err,
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueMessageBind(queue *messages.Queue, bindingKeys []string) error {
	var errWrap error

	for _, bindingKey := range bindingKeys {
		rmq.Logger.Info(fmt.Sprintf("binding queue [%s] to exchange [%s] with binding key [%s]", queue.Name, rmq.exchange, bindingKey))

		err := rmq.Channel.QueueBind(
			queue.Name,
			bindingKey,
			rmq.exchange,
			false,
			nil)

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "AddQueueMessageBind",
				What:  "failure to bind a queue",
				Why: errors.Meta{
					"Queue":       queue.Name,
					"Binding Key": bindingKey,
					"Exchange":    rmq.exchange,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)
		}
	}

	if errWrap != nil {
		return errors.BubbleUp(errWrap, "AddQueueMessageBind")
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueConsumer(consumer messages.Consumer) error {
	var errWrap error

	for _, queue := range consumer.SubscribedTo() {
		deliveries, err := rmq.Channel.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "AddQueueConsumer",
				What:  "failure to register a consumer",
				Why: errors.Meta{
					"Queue":    queue.Name,
					"Exchange": rmq.exchange,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)

			continue
		}

		go func() {
			for delivery := range deliveries {
				message := new(messages.Message)

				err := json.Unmarshal(delivery.Body, message)

				if err != nil {
					rmq.Logger.Error(fmt.Sprintf("failed to deliver a message with id [%s] from queue [%s]", message.Id, queue.Name))
					continue
				}

				err = consumer.On(message)

				if err != nil {
					rmq.Logger.Error(fmt.Sprintf("failed to consume a message with id [%s] from queue [%s]", message.Id, queue.Name))
					continue
				}

				delivery.Ack(false)
			}
		}()
	}

	if errWrap != nil {
		return errors.BubbleUp(errWrap, "AddQueueConsumer")
	}

	return nil
}

func (rmq *RabbitMQ) PublishMessages(messages []*messages.Message) error {
	var errWrap error

	for _, message := range messages {
		if message.Id == "" {
			message.Id = uuid.NewString()
		}

		if message.OccurredOn == "" {
			message.OccurredOn = time.Now().UTC().Format(time.RFC3339Nano)
		}

		body, err := json.Marshal(message)

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "PublishMessages",
				What:  "cannot encode message to json",
				Why: errors.Meta{
					"Exchange": rmq.exchange,
					"Message":  message.Id,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)

			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = rmq.Channel.PublishWithContext(ctx,
			rmq.exchange,
			message.Type,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			})

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "PublishMessages",
				What:  "failure to publish a message",
				Why: errors.Meta{
					"Exchange": rmq.exchange,
					"Message":  message.Id,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)
		}
	}

	if errWrap != nil {
		return errors.BubbleUp(errWrap, "PublishMessages")
	}

	return nil
}

func CloseRabbitMQ(rmq *RabbitMQ) error {
	err := rmq.Channel.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failure to close channel",
			Who:   err,
		})
	}

	err = rmq.Connection.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failure to close rabbitmq connection",
			Who:   err,
		})
	}

	return nil
}

func NewRabbitMQ(uri string, logger models.Logger) (messages.Broker, error) {
	conn, err := amqp.Dial(uri)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewRabbitMQ",
			What:  "failure connecting to rabbitmq",
			Who:   err,
		})
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewRabbitMQ",
			What:  "failure to open a channel",
			Who:   err,
		})
	}

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
		Logger:     logger,
	}, nil
}
