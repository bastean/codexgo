package mongodb

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
)

type MongoDB = *mongodb.MongoDB

func Open(uri, name string) (*mongodb.MongoDB, error) {
	session, err := mongodb.Open(uri, name)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	return session, nil
}

func Close(ctx context.Context, session *mongodb.MongoDB) error {
	err := mongodb.Close(ctx, session)

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
