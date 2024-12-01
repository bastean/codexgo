package sqlite

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

const (
	ErrUniqueConstraintFailed = "UNIQUE constraint failed"
)

const (
	InMemory = "file::memory:?cache=shared"
)

type Database struct {
	Session *gorm.DB
}

func Open(dsn string) (*Database, error) {
	if dsn == "" {
		dsn = InMemory
	}

	session, err := gorm.Open(sqlite.Open(dsn))

	if err != nil {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Open",
			What:  "Failure opening SQLite DSN",
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
	session, err := database.Session.DB()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Close",
			What:  "Failure to obtain database",
			Who:   err,
		})
	}

	err = session.Close()

	if err != nil {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Close",
			What:  "Failure to close database",
			Who:   err,
		})
	}

	return nil
}

func IsErrDuplicateValue(err error) bool {
	if err != nil {
		return strings.Contains(err.Error(), ErrUniqueConstraintFailed)
	}

	return false
}

func HandleErrDuplicateValue(err error) error {
	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.Split(err.Error(), ".")[1])

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
	if err != nil {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}

	return false
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
