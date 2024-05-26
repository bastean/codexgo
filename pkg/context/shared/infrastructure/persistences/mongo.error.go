package persistences

import (
	"regexp"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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
