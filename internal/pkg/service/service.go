package service

import (
	"context"

	"github.com/bastean/codexgo/internal/pkg/service/communication"
	"github.com/bastean/codexgo/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/internal/pkg/service/persistence/mongodb"
	"github.com/bastean/codexgo/internal/pkg/service/transport/smtp"
	"github.com/bastean/codexgo/internal/pkg/service/user"
)

var (
	err      error
	SMTP     smtp.SMTP
	RabbitMQ communication.Broker
	MongoDB  mongodb.MongoDB
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
	log.EstablishingConnectionWith("smtp")

	OpenSMTP()

	log.ConnectionEstablishedWith("smtp")

	log.EstablishingConnectionWith("rabbitmq")

	err = OpenRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith("rabbitmq")

	log.EstablishingConnectionWith("mongodb")

	err = OpenMongoDB()

	if err != nil {
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith("mongodb")

	log.StartingModule("user")

	err = StartUser()

	if err != nil {
		return errors.BubbleUp(err, "Up")
	}

	log.StartedModule("user")

	return nil
}

func CloseRabbitMQ() error {
	err = rabbitmq.Close(RabbitMQ)

	if err != nil {
		return errors.BubbleUp(err, "CloseRabbitMQ")
	}

	return nil
}

func CloseMongoDB(ctx context.Context) error {
	err = mongodb.Close(ctx, MongoDB)

	if err != nil {
		return errors.BubbleUp(err, "CloseMongoDB")
	}

	return nil
}

func Down(ctx context.Context) error {
	log.ClosingConnectionWith("rabbitmq")

	err = CloseRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith("rabbitmq")

	log.ClosingConnectionWith("mongodb")

	err = CloseMongoDB(ctx)

	if err != nil {
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith("mongodb")

	return nil
}
