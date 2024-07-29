package mongodb

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type MongoDB struct {
	*mongo.Client
	*mongo.Database
}

func Open(uri, name string) (*MongoDB, error) {
	options := options.Client().ApplyURI(uri)

	session, err := mongo.Connect(context.Background(), options)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure to create a MongoDB client",
			Who:   err,
		})
	}

	err = session.Ping(context.Background(), nil)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Open",
			What:  "Failure connecting to MongoDB",
			Who:   err,
		})
	}

	return &MongoDB{
		Client:   session,
		Database: session.Database(name),
	}, nil
}

func Close(ctx context.Context, session *MongoDB) error {
	if err := session.Client.Disconnect(ctx); err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close connection with MongoDB",
			Who:   err,
		})
	}

	return nil
}

func HandleDuplicateKeyError(err error) error {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(err.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	return errors.NewAlreadyExist(&errors.Bubble{
		Where: "HandleDuplicateKeyError",
		What:  fmt.Sprintf("%s already registered", field),
		Why: errors.Meta{
			"Field": field,
		},
		Who: err,
	})
}

func HandleDocumentNotFound(index string, err error) error {
	return errors.NewNotExist(&errors.Bubble{
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", index),
		Why: errors.Meta{
			"Index": index,
		},
		Who: err,
	})
}
