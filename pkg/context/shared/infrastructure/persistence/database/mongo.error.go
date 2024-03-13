package database

import (
	"regexp"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func HandleMongoDuplicateKeyError(error error) {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(error.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	panic(errors.AlreadyExist{Message: field + " already registered"})
}

func HandleMongoDocumentNotFound(index string) {
	panic(errors.NotExist{Message: "Not Found " + index})
}
