package badgerdb

import (
	"strings"

	"github.com/dgraph-io/badger/v4"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

const (
	Separator = ":"
)

type Database struct {
	Session *badger.DB
}

func Open(dsn string) (*Database, error) {
	options := badger.DefaultOptions(dsn)

	if dsn == "" {
		options = options.WithInMemory(true)
	}

	session, err := badger.Open(options)

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure opening BadgerDB DSN",
			Why: errors.Meta{
				"DSN": dsn,
			},
			Who: err,
		})
	}

	return &Database{
		Session: session,
	}, nil
}

func Close(database *Database) error {
	if err := database.Session.Close(); err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to close BadgerDB",
			Who:  err,
		})
	}

	return nil
}

func NewKey(values ...string) ([]byte, error) {
	for _, value := range values {
		if strings.TrimSpace(value) == "" || strings.Contains(value, Separator) {
			return nil, errors.New[errors.Internal](&errors.Bubble{
				What: "Cannot create key with invalid values",
				Why: errors.Meta{
					"Value": value,
				},
			})
		}
	}

	return []byte(strings.Join(values, Separator)), nil
}

func ParseKey(key []byte) []string {
	return strings.Split(string(key), Separator)
}

func HandleErrDuplicateValue(field string) error {
	if field == "Id" {
		field = strings.ToUpper(field)
	}

	return errors.New[errors.AlreadyExist](&errors.Bubble{
		What: field + " already registered",
		Why: errors.Meta{
			"Field": field,
		},
	})
}

func IsErrNotFound(err error) bool {
	return err != nil && errors.Is(err, badger.ErrKeyNotFound)
}

func HandleErrNotFound(index string) error {
	return errors.New[errors.NotExist](&errors.Bubble{
		What: index + " not found",
		Why: errors.Meta{
			"Index": index,
		},
	})
}
