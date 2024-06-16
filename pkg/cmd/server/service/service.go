package service

import (
	"context"

	"github.com/bastean/codexgo/pkg/cmd/server/service/broker/rabbitmq"
	"github.com/bastean/codexgo/pkg/cmd/server/service/database/mongodb"
	"github.com/bastean/codexgo/pkg/cmd/server/service/env"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/service/smtp"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
)

var (
	err      error
	SMTP     *transports.SMTP
	RabbitMQ models.Broker
	MongoDB  *persistences.MongoDB
)

func Start() error {
	logger.Info("starting smtp")

	SMTP = smtp.New(
		env.SMTP.Host,
		env.SMTP.Port,
		env.SMTP.Username,
		env.SMTP.Password,
		env.SMTP.ServerURL,
	)

	if env.SMTP.Host != "" {
		user.InitCreated(user.NewMailConfirmation(SMTP))
	} else {
		user.InitCreated(user.NewTerminalConfirmation(
			logger.Logger,
			env.ServerURL,
		))
	}

	logger.Info("starting rabbitmq")

	RabbitMQ, err = rabbitmq.New(
		env.Broker.URI,
		logger.Logger,
		rabbitmq.Exchange,
		rabbitmq.Queues,
		rabbitmq.Consumers(
			user.Created,
		),
	)

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	logger.Info("starting mongodb")

	MongoDB, err = mongodb.New(
		env.Database.URI,
		"codexgo",
	)

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	logger.Info("starting user")

	userMongoCollection, err := user.NewMongoCollection(
		MongoDB,
		"users",
		user.Bcrypt,
	)

	if err != nil {
		return errors.BubbleUp(err, "Start")
	}

	user.Init(
		userMongoCollection,
		RabbitMQ,
		user.Bcrypt,
	)

	return nil
}

func Stop(ctx context.Context) error {
	logger.Info("closing rabbitmq")

	err = rabbitmq.Close(RabbitMQ)

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	logger.Info("closing mongodb")

	err = mongodb.Close(MongoDB, ctx)

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
