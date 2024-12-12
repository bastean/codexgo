package mongodb

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

const (
	ErrNoDocuments = "no documents"
)

type Database struct {
	*mongo.Client
	*mongo.Database
}

func Open(uri, name string) (*Database, error) {
	options := options.Client().ApplyURI(uri)

	session, err := mongo.Connect(context.Background(), options)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure to create a MongoDB client",
			Who:   err,
		})
	}

	err = session.Ping(context.Background(), nil)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure connecting to MongoDB",
			Who:   err,
		})
	}

	return &Database{
		Client:   session,
		Database: session.Database(name),
	}, nil
}

func Close(ctx context.Context, session *Database) error {
	if err := session.Client.Disconnect(ctx); err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Close",
			What:  "Failure to close connection with MongoDB",
			Who:   err,
		})
	}

	return nil
}

func IsErrDuplicateValue(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

func HandleErrDuplicateValue(err error) error {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(err.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	switch field {
	case "Id":
		field = strings.ToUpper(field)
	}

	return errors.New[errors.AlreadyExist](&errors.Bubble{
		Where: "HandleErrDuplicateValue",
		What:  fmt.Sprintf("%s already registered", field),
		Why: errors.Meta{
			"Field": field,
		},
		Who: err,
	})
}

func IsErrNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), ErrNoDocuments)
}

func HandleErrNotFound(err error, index string) error {
	return errors.New[errors.NotExist](&errors.Bubble{
		Where: "HandleErrNotFound",
		What:  fmt.Sprintf("%s not found", index),
		Why: errors.Meta{
			"Index": index,
		},
		Who: err,
	})
}
