package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	loggers.Logger
	exchange string
}

func (rabbitMQ *RabbitMQ) AddRouter(router *messages.Router) error {
	err := rabbitMQ.Channel.ExchangeDeclare(
		router.Name,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "AddRouter",
			What:  "Failure to declare a router",
			Why: errors.Meta{
				"Router": router.Name,
			},
			Who: err,
		})
	}

	rabbitMQ.exchange = router.Name

	return nil
}

func (rabbitMQ *RabbitMQ) AddQueue(queue *messages.Queue) error {
	_, err := rabbitMQ.Channel.QueueDeclare(
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
			What:  "Failure to declare a queue",
			Why: errors.Meta{
				"Queue": queue.Name,
			},
			Who: err,
		})
	}

	return nil
}

func (rabbitMQ *RabbitMQ) AddQueueMessageBind(queue *messages.Queue, bindingKeys []string) error {
	var errWrap error

	for _, bindingKey := range bindingKeys {
		rabbitMQ.Logger.Info(fmt.Sprintf("binding queue [%s] to exchange [%s] with binding key [%s]", queue.Name, rabbitMQ.exchange, bindingKey))

		err := rabbitMQ.Channel.QueueBind(
			queue.Name,
			bindingKey,
			rabbitMQ.exchange,
			false,
			nil,
		)

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "AddQueueMessageBind",
				What:  "Failure to bind a queue",
				Why: errors.Meta{
					"Queue":       queue.Name,
					"Binding Key": bindingKey,
					"Exchange":    rabbitMQ.exchange,
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

func (rabbitMQ *RabbitMQ) AddQueueConsumer(consumer messages.Consumer) error {
	var errWrap error

	for _, queue := range consumer.SubscribedTo() {
		deliveries, err := rabbitMQ.Channel.Consume(
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
				What:  "Failure to register a consumer",
				Why: errors.Meta{
					"Queue":    queue.Name,
					"Exchange": rabbitMQ.exchange,
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
					rabbitMQ.Logger.Error(fmt.Sprintf("failed to deliver a message with id [%s] from queue [%s]", message.Id, queue.Name))
					continue
				}

				err = consumer.On(message)

				if err != nil {
					rabbitMQ.Logger.Error(fmt.Sprintf("failed to consume a message with id [%s] from queue [%s]", message.Id, queue.Name))
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

func (rabbitMQ *RabbitMQ) PublishMessages(messages []*messages.Message) error {
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
				What:  "Cannot encode message to JSON",
				Why: errors.Meta{
					"Exchange": rabbitMQ.exchange,
					"Message":  message.Id,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)

			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = rabbitMQ.Channel.PublishWithContext(ctx,
			rabbitMQ.exchange,
			message.Type,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			},
		)

		if err != nil {
			errToWrap := errors.NewInternal(&errors.Bubble{
				Where: "PublishMessages",
				What:  "Failure to publish a message",
				Why: errors.Meta{
					"Exchange": rabbitMQ.exchange,
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

func Open(uri string, logger loggers.Logger) (*RabbitMQ, error) {
	session, err := amqp.Dial(uri)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure connecting to RabbitMQ",
			Who:   err,
		})
	}

	channel, err := session.Channel()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure to open a channel",
			Who:   err,
		})
	}

	return &RabbitMQ{
		Connection: session,
		Channel:    channel,
		Logger:     logger,
	}, nil
}

func Close(session *RabbitMQ) error {
	err := session.Channel.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close channel",
			Who:   err,
		})
	}

	err = session.Connection.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close RabbitMQ connection",
			Who:   err,
		})
	}

	return nil
}
