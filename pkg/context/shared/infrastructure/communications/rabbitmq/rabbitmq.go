package rabbitmq

import (
	"context"
	"fmt"
	"reflect"

	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/id"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type Queue struct {
	Name             *messages.Recipient
	BindingKey       string
	Attributes, Meta reflect.Type
}

type (
	Mapper map[*messages.Key][]*Queue
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	roles.Logger
	ConsumeCycle context.Context
	exchange     string
	queues       Mapper
}

func (r *RabbitMQ) AddExchange(name string) error {
	err := r.Channel.ExchangeDeclare(
		name,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to declare a Exchange",
			Why: errors.Meta{
				"Exchange": name,
			},
			Who: err,
		})
	}

	r.exchange = name

	return nil
}

func (r *RabbitMQ) AddQueue(name *messages.Recipient) error {
	_, err := r.Channel.QueueDeclare(
		name.Value(),
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to declare a Queue",
			Why: errors.Meta{
				"Queue": name.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (r *RabbitMQ) AddQueueEventBind(queue *Queue, routingKey *messages.Key) error {
	err := r.Channel.QueueBind(
		queue.Name.Value(),
		queue.BindingKey,
		r.exchange,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to bind a Queue",
			Why: errors.Meta{
				"Exchange":    r.exchange,
				"Queue":       queue.Name.Value(),
				"Binding Key": queue.BindingKey,
				"Routing Key": routingKey.Value(),
			},
			Who: err,
		})
	}

	r.Logger.Info(fmt.Sprintf("Binding Queue [%s] to Exchange [%s] with Binding Key [%s]", queue.Name.Value(), r.exchange, queue.BindingKey))

	r.queues[routingKey] = append(r.queues[routingKey], queue)

	return nil
}

func (r *RabbitMQ) Consume(key *messages.Key, queue *Queue, deliveries <-chan amqp.Delivery, consumer roles.EventConsumer) {
	for delivery := range deliveries {
		event := new(messages.Message)

		if queue.Attributes != nil {
			event.Attributes = reflect.New(queue.Attributes.Elem()).Interface()
		}

		if queue.Meta != nil {
			event.Meta = reflect.New(queue.Meta.Elem()).Interface()
		}

		err := json.Unmarshal(delivery.Body, event)

		if err != nil {
			r.Logger.Error(fmt.Sprintf("Failure to decode an Event with ID [%s] from Queue [%s]: [%s]", key.Value(), queue.Name.Value(), err))
			continue
		}

		event.Key = key

		err = consumer.On(event)

		if err != nil {
			r.Logger.Error(fmt.Sprintf("Failure to consume an Event with ID [%s] from Queue [%s]: [%s]", event.Key.Value(), queue.Name.Value(), err))
			continue
		}

		err = delivery.Ack(false)

		if err != nil {
			r.Logger.Error(fmt.Sprintf("Failure to deliver an acknowledgement for Event with ID [%s] to Queue [%s]: [%s]", event.Key.Value(), queue.Name.Value(), err))
		}
	}
}

func (r *RabbitMQ) Subscribe(key *messages.Key, consumer roles.EventConsumer) error {
	queues, ok := r.queues[key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Queue is not declared",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Event":    key.Value(),
			},
		})
	}

	for _, queue := range queues {
		deliveries, err := r.Channel.ConsumeWithContext(
			r.ConsumeCycle,
			queue.Name.Value(),
			"",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			return errors.New[errors.Internal](&errors.Bubble{
				What: "Failure to subscribe a Consumer",
				Why: errors.Meta{
					"Exchange": r.exchange,
					"Queue":    queue.Name.Value(),
				},
				Who: err,
			})
		}

		go r.Consume(key, queue, deliveries, consumer)
	}

	return nil
}

func (r *RabbitMQ) Publish(event *messages.Message) error {
	_, ok := r.queues[event.Key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to execute a Event without a Consumer",
			Why: errors.Meta{
				"Event": event.Key.Value(),
			},
		})
	}

	if event.ID == "" {
		event.ID = id.New()
	}

	if event.OccurredOn == "" {
		event.OccurredOn = time.Now().Format()
	}

	body, err := json.Marshal(event)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Cannot encode Event to JSON",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Event":    event.Key.Value(),
			},
			Who: err,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = r.Channel.PublishWithContext(
		ctx,
		r.exchange,
		event.Key.Value(),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to publish a Event",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Event":    event.Key.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func Open(uri string, exchange string, logger roles.Logger, consumeCycle context.Context) (*RabbitMQ, error) {
	session, err := amqp.Dial(uri)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure connecting to RabbitMQ",
			Who:  err,
		})
	}

	channel, err := session.Channel()

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to open a Channel",
			Who:  err,
		})
	}

	rmq := &RabbitMQ{
		Connection:   session,
		Channel:      channel,
		Logger:       logger,
		ConsumeCycle: consumeCycle,
		queues:       make(Mapper),
	}

	err = rmq.AddExchange(exchange)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return rmq, nil
}

func Close(session *RabbitMQ) error {
	err := session.Channel.Close()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to close Channel",
			Who:  err,
		})
	}

	err = session.Connection.Close()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to close RabbitMQ connection",
			Who:  err,
		})
	}

	return nil
}

func AddQueueMapper(rmq *RabbitMQ, mapper Mapper) error {
	var err error

	for routingKey, queues := range mapper {
		for _, queue := range queues {
			err = rmq.AddQueue(queue.Name)

			if err != nil {
				return errors.BubbleUp(err)
			}

			err = rmq.AddQueueEventBind(queue, routingKey)

			if err != nil {
				return errors.BubbleUp(err)
			}
		}
	}

	return nil
}
