package persistences

import (
	"context"
	"regexp"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type MongoDB struct {
	*mongo.Client
	*mongo.Database
}

func NewMongoDatabase(uri, databaseName string) (*MongoDB, error) {
	var err error

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMongoDatabase",
			What:  "failure to create a mongodb client",
			Who:   err,
		})
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
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

func HandleMongoDuplicateKeyError(err error) error {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(err.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	return errors.NewAlreadyExist(&errors.Bubble{
		Where: "HandleMongoDuplicateKeyError",
		What:  "already registered",
		Why: errors.Meta{
			"Field": field,
		},
		Who: err,
	})
}

func HandleMongoDocumentNotFound(index string, err error) error {
	return errors.NewNotExist(&errors.Bubble{
		Where: "HandleMongoDocumentNotFound",
		What:  "not found",
		Why: errors.Meta{
			"Index": index,
		},
		Who: err,
	})
}

func CloseMongoDatabase(ctx context.Context, mdb *MongoDB) error {
	err := mdb.Client.Disconnect(ctx)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseMongoDatabase",
			What:  "failure to close connection with mongodb",
			Who:   err,
		})
	}

	return nil
}
