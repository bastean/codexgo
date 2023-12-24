package persistence

import (
	"regexp"
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
)

func handleDuplicateKeyError(error error) {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(error.Error())

	field := strings.Title(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	panic(errors.AlreadyExist{Dump: []errors.Error{{Field: field, Message: "Duplicate"}}})
}

func handleDocumentNotFound(index string) {
	panic(errors.NotExist{Dump: []errors.Error{{Field: index, Message: "Not Found"}}})
}
