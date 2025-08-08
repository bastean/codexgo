package embed

import (
	"regexp"
)

const (
	RExEmbed = `\[.*]`
)

var (
	RExEmbedDo = regexp.MustCompile(RExEmbed)
)
