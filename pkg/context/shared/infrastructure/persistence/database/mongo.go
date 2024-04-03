package database

import (
	"context"

	"github.com/bastean/codexgo/pkg/context/shared/domain/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Client
	*mongo.Database
}

func CloseMongoDatabase(ctx context.Context, mdb *MongoDB) {
	err := mdb.Client.Disconnect(ctx)

	if err != nil {
		service.FailOnError(err, "failed to close mongodb connection")
	}
}

func NewMongoDatabase(uri, databaseName string) *MongoDB {
	var err error

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	return &MongoDB{
		Client:   client,
		Database: client.Database(databaseName)}
}
