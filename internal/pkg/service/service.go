package service

import (
	"context"

	"github.com/bastean/codexgo/internal/pkg/service/communication"
	"github.com/bastean/codexgo/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/internal/pkg/service/env"
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/internal/pkg/service/logger"
	"github.com/bastean/codexgo/internal/pkg/service/persistence/mongodb"
	"github.com/bastean/codexgo/internal/pkg/service/transport/smtp"
	"github.com/bastean/codexgo/internal/pkg/service/user"
)

var (
	err      error
	RabbitMQ communication.Broker
	MongoDB  mongodb.MongoDB
	SMTP     smtp.SMTP
)

func OpenSMTP() {
	if env.SMTP.Host == "" {
		user.InitCreated(user.TerminalConfirmation(logger.Logger, env.ServerURL), user.QueueSendConfirmation)
		return
	}

	SMTP = smtp.New(
		env.SMTP.Host,
		env.SMTP.Port,
		env.SMTP.Username,
		env.SMTP.Password,
		env.SMTP.ServerURL,
	)

	user.InitCreated(user.MailConfirmation(SMTP), user.QueueSendConfirmation)
}

func OpenRabbitMQ() error {
	RabbitMQ, err = rabbitmq.New(
		env.RabbitMQ.URI,
		logger.Logger,
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
	MongoDB, err = mongodb.New(
		env.Mongo.URI,
		env.Mongo.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenMongoDB")
	}

	return nil
}

func StartUser() error {
	collection, err := user.MongoCollection(
		MongoDB,
		"users",
		user.Bcrypt,
	)

	if err != nil {
		return errors.BubbleUp(err, "StartUser")
	}

	user.Init(
		collection,
		RabbitMQ,
		user.Bcrypt,
	)

	return nil
}

func Run() error {
	logger.EstablishingConnectionWith("smtp")

	OpenSMTP()

	logger.ConnectionEstablishedWith("smtp")

	logger.EstablishingConnectionWith("rabbitmq")

	err = OpenRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	logger.ConnectionEstablishedWith("rabbitmq")

	logger.EstablishingConnectionWith("mongodb")

	err = OpenMongoDB()

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	logger.ConnectionEstablishedWith("mongodb")

	logger.StartingModule("user")

	err = StartUser()

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	logger.StartedModule("user")

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
	err = mongodb.Close(MongoDB, ctx)

	if err != nil {
		return errors.BubbleUp(err, "CloseMongoDB")
	}

	return nil
}

func Stop(ctx context.Context) error {
	logger.ClosingConnectionWith("rabbitmq")

	err = CloseRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	logger.ConnectionClosedWith("rabbitmq")

	logger.ClosingConnectionWith("mongodb")

	err = CloseMongoDB(ctx)

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	logger.ConnectionClosedWith("mongodb")

	return nil
}
