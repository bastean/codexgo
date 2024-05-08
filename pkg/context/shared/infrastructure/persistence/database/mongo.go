package database

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Client
	*mongo.Database
}

func CloseMongoDatabase(ctx context.Context, mdb *MongoDB) error {
	err := mdb.Client.Disconnect(ctx)

	if err != nil {
		return errs.NewInternalError(&errs.Bubble{
			Where: "CloseMongoDatabase",
			What:  "failed to close mongodb connection",
			Who:   err,
		})
	}

	return nil
}

func NewMongoDatabase(uri, databaseName string) (*MongoDB, error) {
	var err error

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, errs.NewInternalError(&errs.Bubble{
			Where: "NewMongoDatabase",
			What:  "failed to create a mongodb client",
			Who:   err,
		})
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, errs.NewInternalError(&errs.Bubble{
			Where: "NewMongoDatabase",
			What:  "failed to connect with mongodb",
			Who:   err,
		})
	}

	return &MongoDB{
		Client:   client,
		Database: client.Database(databaseName),
	}, nil
}
