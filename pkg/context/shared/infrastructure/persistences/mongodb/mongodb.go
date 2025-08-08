package mongodb

import (
	"context"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/array"
)

const (
	ErrNoDocuments = "no documents"
)

const (
	RExDuplicateValue = `[A-Za-z0-9]+\.value`
)

var (
	RExDuplicateValueDo = regexp.MustCompile(RExDuplicateValue)
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
			What: "Failure to create a MongoDB client",
			Who:  err,
		})
	}

	err = session.Ping(context.Background(), nil)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure connecting to MongoDB",
			Who:  err,
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
			What: "Failure to close connection with MongoDB",
			Who:  err,
		})
	}

	return nil
}

func IsErrDuplicateValue(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

func HandleErrDuplicateValue(err error) error {
	field, exists := array.Slice(strings.Split(RExDuplicateValueDo.FindString(err.Error()), "."), 0)

	if !exists {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Missing field",
			Who:  err,
		})
	}

	field = cases.Title(language.English).String(field)

	if field == "Id" {
		field = strings.ToUpper(field)
	}

	return errors.New[errors.AlreadyExist](&errors.Bubble{
		What: field + " already registered",
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
		What: index + " not found",
		Why: errors.Meta{
			"Index": index,
		},
		Who: err,
	})
}
