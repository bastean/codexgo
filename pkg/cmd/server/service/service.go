package service

import (
	"context"

	"github.com/bastean/codexgo/pkg/cmd/server/service/communication"
	"github.com/bastean/codexgo/pkg/cmd/server/service/communication/rabbitmq"
	"github.com/bastean/codexgo/pkg/cmd/server/service/env"
	"github.com/bastean/codexgo/pkg/cmd/server/service/errors"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/service/persistence/mongodb"
	"github.com/bastean/codexgo/pkg/cmd/server/service/transport/smtp"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
)

var (
	err      error
	RabbitMQ communication.Broker
	MongoDB  mongodb.MongoDB
	SMTP     smtp.SMTP
)

func startSMTP() {
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

func startRabbitMQ() error {
	RabbitMQ, err = rabbitmq.New(
		env.Broker.URI,
		logger.Logger,
		rabbitmq.Exchange("codexgo"),
		rabbitmq.Queues{
			user.QueueSendConfirmation,
		},
		rabbitmq.Consumers{
			user.Created,
		},
	)

	if err != nil {
		return errors.BubbleUp(err, "startRabbitMQ")
	}

	return nil
}

func startMongoDB() error {
	MongoDB, err = mongodb.New(
		env.Database.URI,
		"codexgo",
	)

	if err != nil {
		return errors.BubbleUp(err, "startMongoDB")
	}

	return nil
}

func startUser() error {
	collection, err := user.MongoCollection(
		MongoDB,
		"users",
		user.Bcrypt,
	)

	if err != nil {
		return errors.BubbleUp(err, "startUser")
	}

	user.Init(
		collection,
		RabbitMQ,
		user.Bcrypt,
	)

	return nil
}

func Start() error {
	logger.Info("starting smtp")

	startSMTP()

	logger.Info("starting rabbitmq")

	err = startRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	logger.Info("starting mongodb")

	err = startMongoDB()

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	logger.Info("starting user")

	err = startUser()

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	return nil
}

func stopRabbitMQ() error {
	err = rabbitmq.Close(RabbitMQ)

	if err != nil {
		return errors.BubbleUp(err, "stopRabbitMQ")
	}

	return nil
}

func stopMongoDB(ctx context.Context) error {
	err = mongodb.Close(MongoDB, ctx)

	if err != nil {
		return errors.BubbleUp(err, "stopMongoDB")
	}

	return nil
}

func Stop(ctx context.Context) error {
	logger.Info("stopping rabbitmq")

	err = stopRabbitMQ()

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	logger.Info("stopping mongodb")

	err = stopMongoDB(ctx)

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
