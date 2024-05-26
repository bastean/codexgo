package spersistence

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
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
		return serror.NewInternal(&serror.Bubble{
			Where: "CloseMongoDatabase",
			What:  "failure to close connection with mongodb",
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
		return nil, serror.NewInternal(&serror.Bubble{
			Where: "NewMongoDatabase",
			What:  "failure to create a mongodb client",
			Who:   err,
		})
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, serror.NewInternal(&serror.Bubble{
			Where: "NewMongoDatabase",
			What:  "failure connecting to mongodb",
			Who:   err,
		})
	}

	return &MongoDB{
		Client:   client,
		Database: client.Database(databaseName),
	}, nil
}
