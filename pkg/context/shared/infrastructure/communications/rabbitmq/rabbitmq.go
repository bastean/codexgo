package rabbitmq

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type Recipient struct {
	Name             messages.Recipient
	BindingKey       messages.Key
	Attributes, Meta reflect.Type
}

type (
	QueueMapper map[messages.Key]*Recipient
)

type RabbitMQ struct {
	*amqp.Connection
	*amqp.Channel
	roles.Logger
	ConsumeCycle context.Context
	exchange     string
	queues       QueueMapper
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
			Where: "AddExchange",
			What:  "Failure to declare a Exchange",
			Why: errors.Meta{
				"Exchange": name,
			},
			Who: err,
		})
	}

	r.exchange = name

	return nil
}

func (r *RabbitMQ) AddQueue(name messages.Recipient) error {
	_, err := r.Channel.QueueDeclare(
		string(name),
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "AddQueue",
			What:  "Failure to declare a Queue",
			Why: errors.Meta{
				"Queue": name,
			},
			Who: err,
		})
	}
	return nil
}

func (r *RabbitMQ) AddQueueEventBind(queue messages.Recipient, bindingKey, routingKey messages.Key, attributes, meta reflect.Type) error {
	err := r.Channel.QueueBind(
		string(queue),
		string(bindingKey),
		r.exchange,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "AddQueueEventBind",
			What:  "Failure to bind a Queue",
			Why: errors.Meta{
				"Exchange":    r.exchange,
				"Queue":       queue,
				"Binding Key": bindingKey,
				"Routing Key": routingKey,
			},
			Who: err,
		})
	}

	r.Logger.Info(fmt.Sprintf("Binding Queue [%s] to Exchange [%s] with Binding Key [%s]", queue, r.exchange, bindingKey))

	r.queues[routingKey] = &Recipient{queue, bindingKey, attributes, meta}

	return nil
}

func (r *RabbitMQ) Unmarshal(data []byte, attributes, meta reflect.Type, event *messages.Message) error {
	received := make(map[string]json.RawMessage)

	err := json.Unmarshal(data, &received)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Unmarshal",
			What:  "Cannot unmarshal an Event",
			Who:   err,
		})
	}

	err = json.Unmarshal(data, event)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Unmarshal",
			What:  "Cannot unmarshal an Event ID, OccurredOn & Key",
			Who:   err,
		})
	}

	var value any

	if attributes != nil {
		value = reflect.New(attributes.Elem()).Interface()

		err = json.Unmarshal(received["Attributes"], value)

		if err != nil {
			return errors.New[errors.Internal](&errors.Bubble{
				Where: "Unmarshal",
				What:  "Cannot unmarshal an Event Attributes",
				Who:   err,
			})
		}

		event.Attributes = value
	}

	if meta != nil {
		value = reflect.New(meta.Elem()).Interface()

		err = json.Unmarshal(received["Meta"], value)

		if err != nil {
			return errors.New[errors.Internal](&errors.Bubble{
				Where: "Unmarshal",
				What:  "Cannot unmarshal an Event Meta",
				Who:   err,
			})
		}

		event.Meta = value
	}

	return nil
}

func (r *RabbitMQ) Consume(key messages.Key, queue *Recipient, deliveries <-chan amqp.Delivery, consumer roles.EventConsumer) {
	for delivery := range deliveries {
		event := new(messages.Message)

		err := r.Unmarshal(delivery.Body, queue.Attributes, queue.Meta, event)

		if err != nil {
			r.Logger.Error(fmt.Sprintf("Failed to deliver a Event with ID [%s] from Queue [%s]: [%s]", key, queue, err))
			continue
		}

		err = consumer.On(event)

		if err != nil {
			r.Logger.Error(fmt.Sprintf("Failed to consume a Event with ID [%s] from Queue [%s]: [%s]", key, queue, err))
			continue
		}

		delivery.Ack(false)
	}
}

func (r *RabbitMQ) Subscribe(key messages.Key, consumer roles.EventConsumer) error {
	queue, ok := r.queues[key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Subscribe",
			What:  "Queue is not declared",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Event":    key,
			},
		})
	}

	deliveries, err := r.Channel.ConsumeWithContext(
		r.ConsumeCycle,
		string(queue.Name),
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Subscribe",
			What:  "Failure to subscribe a Consumer",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Queue":    queue,
			},
			Who: err,
		})
	}

	go r.Consume(key, queue, deliveries, consumer)

	return nil
}

func (r *RabbitMQ) Publish(event *messages.Message) error {
	queue, ok := r.queues[event.Key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Publish",
			What:  "Failure to execute a Event without a Consumer",
			Why: errors.Meta{
				"Event": event.Key,
			},
		})
	}

	if event.ID == "" {
		event.ID = services.GenerateID()
	}

	if event.OccurredOn == "" {
		event.OccurredOn = services.TimeNow()
	}

	body, err := json.Marshal(event)

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Publish",
			What:  "Cannot encode Event to JSON",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Queue":    queue,
				"Event":    event,
			},
			Who: err,
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = r.Channel.PublishWithContext(
		ctx,
		r.exchange,
		string(event.Key),
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
			Where: "Publish",
			What:  "Failure to publish a Event",
			Why: errors.Meta{
				"Exchange": r.exchange,
				"Queue":    queue,
				"Event":    event,
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
			Where: "Open",
			What:  "Failure connecting to RabbitMQ",
			Who:   err,
		})
	}

	channel, err := session.Channel()

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure to open a Channel",
			Who:   err,
		})
	}

	rmq := &RabbitMQ{
		Connection:   session,
		Channel:      channel,
		Logger:       logger,
		ConsumeCycle: consumeCycle,
		queues:       make(QueueMapper),
	}

	err = rmq.AddExchange(exchange)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	return rmq, nil
}

func Close(session *RabbitMQ) error {
	err := session.Channel.Close()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Close",
			What:  "Failure to close Channel",
			Who:   err,
		})
	}

	err = session.Connection.Close()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Close",
			What:  "Failure to close RabbitMQ connection",
			Who:   err,
		})
	}

	return nil
}

func AddQueueMapper(rmq *RabbitMQ, queues QueueMapper) error {
	var err error

	for routingKey, queue := range queues {
		err = rmq.AddQueue(queue.Name)

		if err != nil {
			return errors.BubbleUp(err, "AddQueueMapper")
		}

		err = rmq.AddQueueEventBind(queue.Name, queue.BindingKey, routingKey, queue.Attributes, queue.Meta)

		if err != nil {
			return errors.BubbleUp(err, "AddQueueMapper")
		}
	}

	return nil
}
