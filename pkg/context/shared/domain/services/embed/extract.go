package embed

import (
	"strings"
)

func Extract(message string) string {
	return strings.Trim(RExEmbedDo.FindString(message), "[]")
}
