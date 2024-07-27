package service

import (
	"context"

	"github.com/bastean/codexgo/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/internal/pkg/service/persistence/mongodb"
	"github.com/bastean/codexgo/internal/pkg/service/transport/smtp"
	"github.com/bastean/codexgo/internal/pkg/service/user"
)

var (
	Service = &struct {
		SMTP, RabbitMQ, MongoDB string
	}{
		SMTP:     log.Service("SMTP"),
		RabbitMQ: log.Service("RabbitMQ"),
		MongoDB:  log.Service("MongoDB"),
	}
	Module = &struct {
		User string
	}{
		User: log.Module("User"),
	}
)

var (
	err      error
	SMTP     *smtp.SMTP
	RabbitMQ *rabbitmq.RabbitMQ
	MongoDB  *mongodb.MongoDB
)

func OpenSMTP() {
	if env.SMTP.Host == "" {
		user.InitCreated(user.TerminalConfirmation(log.Log, env.Server.URL), user.QueueSendConfirmation)
		return
	}

	SMTP = smtp.Open(
		env.SMTP.Host,
		env.SMTP.Port,
		env.SMTP.Username,
		env.SMTP.Password,
		env.Server.URL,
	)

	user.InitCreated(user.MailConfirmation(SMTP), user.QueueSendConfirmation)
}

func OpenRabbitMQ() error {
	RabbitMQ, err = rabbitmq.Open(
		env.RabbitMQ.URI,
		log.Log,
		rabbitmq.Exchange(env.RabbitMQ.Name),
		rabbitmq.Queues{
			user.QueueSendConfirmation,
		},
		rabbitmq.Consumers{
			user.Created,
		},
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenRabbitMQ")
	}

	return nil
}

func OpenMongoDB() error {
	MongoDB, err = mongodb.Open(
		env.MongoDB.URI,
		env.MongoDB.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenMongoDB")
	}

	return nil
}

func StartUser() error {
	collection, err := user.Collection(
		MongoDB,
		"users",
		user.Bcrypt,
	)

	if err != nil {
		return errors.BubbleUp(err, "StartUser")
	}

	user.Start(
		collection,
		RabbitMQ,
		user.Bcrypt,
	)

	return nil
}

func Up() error {
	log.EstablishingConnectionWith(Service.SMTP)

	OpenSMTP()

	log.ConnectionEstablishedWith(Service.SMTP)

	log.EstablishingConnectionWith(Service.RabbitMQ)

	if err = OpenRabbitMQ(); err != nil {
		log.ConnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.RabbitMQ)

	log.EstablishingConnectionWith(Service.MongoDB)

	if err = OpenMongoDB(); err != nil {
		log.ConnectionFailedWith(Service.MongoDB)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.MongoDB)

	log.Starting(Module.User)

	if err = StartUser(); err != nil {
		log.CannotBeStarted(Module.User)
		return errors.BubbleUp(err, "Up")
	}

	log.Started(Module.User)

	return nil
}

func CloseRabbitMQ() error {
	if err = rabbitmq.Close(RabbitMQ); err != nil {
		return errors.BubbleUp(err, "CloseRabbitMQ")
	}

	return nil
}

func CloseMongoDB(ctx context.Context) error {
	if err = mongodb.Close(ctx, MongoDB); err != nil {
		return errors.BubbleUp(err, "CloseMongoDB")
	}

	return nil
}

func Down(ctx context.Context) error {
	log.ClosingConnectionWith(Service.RabbitMQ)

	if err = CloseRabbitMQ(); err != nil {
		log.DisconnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.RabbitMQ)

	log.ClosingConnectionWith(Service.MongoDB)

	if err = CloseMongoDB(ctx); err != nil {
		log.DisconnectionFailedWith(Service.MongoDB)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.MongoDB)

	return nil
}
