package mongodb

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
)

type MongoDB = *mongodb.MongoDB

func New(uri, name string) (*mongodb.MongoDB, error) {
	mongoDB, err := mongodb.New(uri, name)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return mongoDB, nil
}

func Close(ctx context.Context, connection *mongodb.MongoDB) error {
	err := mongodb.Close(ctx, connection)

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
