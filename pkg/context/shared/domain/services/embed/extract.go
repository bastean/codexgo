package embed

import (
	"regexp"
	"strings"
)

func Extract(message string) string {
	return strings.Trim(regexp.MustCompile(RExEmbed).FindString(message), "[]")
}
