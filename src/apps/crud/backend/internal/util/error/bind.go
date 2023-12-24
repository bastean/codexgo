package error

import (
	"regexp"
	"strings"
)

func Bind(err string) (error []Error) {
	re := regexp.MustCompile(`for '[A-za-z0-9]+'`)

	rawFields := re.FindAllString(err, -1)

	error = make([]Error, len(rawFields))

	for i, rawField := range rawFields {
		re = regexp.MustCompile(`'[A-Za-z0-9]+'`)

		field := strings.Trim(re.FindString(rawField), "'")
		message := "Required"

		error[i] = Error{field, message}
	}

	return
}
