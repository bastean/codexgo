package service

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence/mongodb"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport/smtp"
	"github.com/bastean/codexgo/v4/internal/pkg/service/user"
)

var Service = &struct {
	SMTP, RabbitMQ, MongoDB, Terminal string
}{
	SMTP:     log.Service("SMTP"),
	RabbitMQ: log.Service("RabbitMQ"),
	MongoDB:  log.Service("MongoDB"),
	Terminal: log.Service("Terminal"),
}

var Module = &struct {
	User string
}{
	User: log.Module("User"),
}

var (
	err      error
	SMTP     *smtp.SMTP
	RabbitMQ *rabbitmq.RabbitMQ
	MongoDB  *mongodb.MongoDB
)

func Up() error {
	switch {
	case env.SMTPHost != "":
		log.EstablishingConnectionWith(Service.SMTP)

		SMTP = smtp.Open(
			&smtp.Auth{
				Host:     env.SMTPHost,
				Port:     env.SMTPPort,
				Username: env.SMTPUsername,
				Password: env.SMTPPassword,
			},
		)

		user.InitCreated(&user.MailConfirmation{
			SMTP:         SMTP,
			AppServerURL: env.ServerGinURL,
		}, user.QueueSendConfirmation)

		log.ConnectionEstablishedWith(Service.SMTP)
	default:
		log.Starting(Service.Terminal)

		user.InitCreated(&user.TerminalConfirmation{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		},
			user.QueueSendConfirmation,
		)

		log.Started(Service.Terminal)
	}

	log.EstablishingConnectionWith(Service.RabbitMQ)

	RabbitMQ, err = rabbitmq.Open(
		env.BrokerRabbitMQURI,
		log.Log,
		rabbitmq.Exchange(env.BrokerRabbitMQName),
		rabbitmq.Queues{
			user.QueueSendConfirmation,
		},
		rabbitmq.Consumers{
			user.Created,
		},
	)

	if err != nil {
		log.ConnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.RabbitMQ)

	log.EstablishingConnectionWith(Service.MongoDB)

	MongoDB, err = mongodb.Open(
		env.DatabaseMongoDBURI,
		env.DatabaseMongoDBName,
	)

	if err != nil {
		log.ConnectionFailedWith(Service.MongoDB)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.MongoDB)

	log.Starting(Module.User)

	collection, err := user.OpenCollection(
		MongoDB,
		"users",
		user.Bcrypt,
	)

	if err != nil {
		return errors.BubbleUp(err, "Up")
	}

	user.Start(
		collection,
		RabbitMQ,
		user.Bcrypt,
	)

	log.Started(Module.User)

	return nil
}

func Down(ctx context.Context) error {
	log.ClosingConnectionWith(Service.RabbitMQ)

	if err = rabbitmq.Close(RabbitMQ); err != nil {
		log.DisconnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.RabbitMQ)

	log.ClosingConnectionWith(Service.MongoDB)

	if err = mongodb.Close(ctx, MongoDB); err != nil {
		log.DisconnectionFailedWith(Service.MongoDB)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.MongoDB)

	return nil
}
