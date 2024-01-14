package persistence

import (
	"regexp"
	"strings"

	"github.com/bastean/codexgo/context/pkg/shared/domain/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func handleDuplicateKeyError(error error) {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(error.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	panic(errors.AlreadyExist{Message: field + " already registered"})
}

func handleDocumentNotFound(index string) {
	panic(errors.NotExist{Message: "Not Found " + index})
}
