package scommunication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/srouter"
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	smodel.Logger
	exchange string
}

func (rmq *RabbitMQ) AddRouter(router *srouter.Router) error {
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
		return serror.NewInternalError(&serror.Bubble{
			Where: "AddRouter",
			What:  "failed to declare an exchange",
			Why: serror.Meta{
				"Router": router.Name,
			},
			Who: err,
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueue(queue *squeue.Queue) error {
	_, err := rmq.Channel.QueueDeclare(
		queue.Name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return serror.NewInternalError(&serror.Bubble{
			Where: "AddQueue",
			What:  "failed to declare a queue",
			Why: serror.Meta{
				"Queue": queue.Name,
			},
			Who: err,
		})
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueMessageBind(queue *squeue.Queue, bindingKeys []string) error {
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
			errToWrap := serror.NewInternalError(&serror.Bubble{
				Where: "AddQueueMessageBind",
				What:  "failed to bind a queue",
				Why: serror.Meta{
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
		return serror.BubbleUp(errWrap, "AddQueueMessageBind")
	}

	return nil
}

func (rmq *RabbitMQ) AddQueueConsumer(consumer smodel.Consumer) error {
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
			errToWrap := serror.NewInternalError(&serror.Bubble{
				Where: "AddQueueConsumer",
				What:  "failed to register a consumer",
				Why: serror.Meta{
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
				message := new(smessage.Message)

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
		return serror.BubbleUp(errWrap, "AddQueueConsumer")
	}

	return nil
}

func (rmq *RabbitMQ) PublishMessages(messages []*smessage.Message) error {
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
			errToWrap := serror.NewInternalError(&serror.Bubble{
				Where: "PublishMessages",
				What:  "failed to publish a messages",
				Why: serror.Meta{
					"Exchange": rmq.exchange,
					"Message":  message.Id,
				},
				Who: err,
			})

			errWrap = errors.Join(errWrap, errToWrap)
		}
	}

	if errWrap != nil {
		return serror.BubbleUp(errWrap, "PublishMessages")
	}

	return nil
}

func CloseRabbitMQ(rmq *RabbitMQ) error {
	err := rmq.Channel.Close()

	if err != nil {
		return serror.NewInternalError(&serror.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failed to close channel",
			Who:   err,
		})
	}

	err = rmq.Connection.Close()

	if err != nil {
		return serror.NewInternalError(&serror.Bubble{
			Where: "CloseRabbitMQ",
			What:  "failed to close rabbitmq connection",
			Who:   err,
		})
	}

	return nil
}

func NewRabbitMQ(uri string, logger smodel.Logger) (smodel.Broker, error) {
	conn, err := amqp.Dial(uri)

	if err != nil {
		return nil, serror.NewInternalError(&serror.Bubble{
			Where: "NewRabbitMQ",
			What:  "failed to connect with rabbitmq",
			Who:   err,
		})
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, serror.NewInternalError(&serror.Bubble{
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
