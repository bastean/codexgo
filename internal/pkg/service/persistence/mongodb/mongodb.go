package mongodb

import (
	"context"

	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
)

type MongoDB = mongodb.MongoDB

func Open(uri, name string) (*MongoDB, error) {
	session, err := mongodb.Open(uri, name)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	return session, nil
}

func Close(ctx context.Context, session *MongoDB) error {
	if err := mongodb.Close(ctx, session); err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
