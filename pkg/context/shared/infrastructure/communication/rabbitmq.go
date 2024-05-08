package communication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	model.Logger
	exchange string
}

func (rmq *RabbitMQ) AddExchange(exchange *exchange.Exchange) error {
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

	if err != nil {
		return errs.NewInternalError(&errs.Bubble{
			Where: "AddExchange",
			What:  "failed to declare an exchange",
			Why: errs.Meta{
				"Exchange": exchange.Name,
			},
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueue(queue *queue.Queue) error {
	_, err := rmq.Channel.QueueDeclare(
		queue.Name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errs.NewInternalError(&errs.Bubble{
			Where: "AddQueue",
			What:  "failed to declare a queue",
			Why: errs.Meta{
				"Queue": queue.Name,
			},
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) error {
	var errWrap error

	for _, bindingKey := range bindingKeys {
		log.Printf("binding queue %s to exchange %s with routing key %s",
			queue.Name, rmq.exchange, bindingKey)

		err := rmq.Channel.QueueBind(
			queue.Name,
			bindingKey,
			rmq.exchange,
			false,
			nil)

		if err != nil {
			errToWrap := errs.NewInternalError(&errs.Bubble{
				Where: "AddQueueMessageBind",
				What:  "failed to bind a queue",
				Why: errs.Meta{
					"Queue":      queue.Name,
					"BindingKey": bindingKey,
					"Exchange":   rmq.exchange,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)
		}
	}

	if errWrap != nil {
		return errs.BubbleUp("AddQueueMessageBind", errWrap)
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueConsumer(consumer model.Consumer) error {
	var errWrap error

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

		if err != nil {
			errToWrap := errs.NewInternalError(&errs.Bubble{
				Where: "AddQueueConsumer",
				What:  "failed to register a consumer",
				Why: errs.Meta{
					"Queue":    queue.Name,
					"Exchange": rmq.exchange,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)

			continue
		}

		go func() {
			for delivery := range messages {
				message := new(message.Message)

				err := json.Unmarshal(delivery.Body, message)

				if err != nil {
					rmq.Logger.Error(fmt.Sprintf("failed to deliver a message with Id:%s from Queue:%s", message.Id, queue.Name))
					continue
				}

				err = consumer.On(message)

				if err != nil {
					rmq.Logger.Error(fmt.Sprintf("failed to consume a message with Id:%s from Queue:%s", message.Id, queue.Name))
					continue
				}

				delivery.Ack(false)
			}
		}()
	}

	if errWrap != nil {
		return errs.BubbleUp("AddQueueConsumer", errWrap)
	}

	return nil
}

func (rmq *RabbitMQ) PublishMessages(messages []*message.Message) error {
	var errWrap error

	for _, message := range messages {
		if message.Id == "" {
			message.Id = uuid.NewString()
		}

		if message.OccurredOn == "" {
			message.OccurredOn = time.Now().UTC().Format(time.DateTime)
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

		if err != nil {
			errToWrap := errs.NewInternalError(&errs.Bubble{
				Where: "PublishMessages",
				What:  "failed to publish a messages",
				Why: errs.Meta{
					"Exchange": rmq.exchange,
					"Message":  message.Id,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)
		}
	}

	if errWrap != nil {
		return errs.BubbleUp("PublishMessages", errWrap)
	}

	return nil
}

func CloseRabbitMQ(rmq *RabbitMQ) error {
	err := rmq.Channel.Close()

	if err != nil {
		return errs.NewInternalError(&errs.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failed to close channel",
			Who:   err,
		})
	}

	err = rmq.Connection.Close()

	if err != nil {
		return errs.NewInternalError(&errs.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failed to close rabbitmq connection",
			Who:   err,
		})
	}

	return nil
}

func NewRabbitMQ(uri string, logger model.Logger) (model.Broker, error) {
	conn, err := amqp.Dial(uri)

	if err != nil {
		return nil, errs.NewInternalError(&errs.Bubble{
			Where: "NewRabbitMQ",
			What:  "failed to connect with rabbitmq",
			Who:   err,
		})
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, errs.NewInternalError(&errs.Bubble{
			Where: "NewRabbitMQ",
			What:  "failed to open a channel",
			Who:   err,
		})
	}

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
		Logger:     logger,
	}, nil
}
