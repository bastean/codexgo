package mongodb

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences"
)

type MongoDB = *persistences.MongoDB

func New(uri, name string) (*persistences.MongoDB, error) {
	mongoDB, err := persistences.NewMongoDatabase(uri, name)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return mongoDB, nil
}

func Close(mongoDB *persistences.MongoDB, ctx context.Context) error {
	err := persistences.CloseMongoDatabase(ctx, mongoDB)

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
