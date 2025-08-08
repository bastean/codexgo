package caller

import (
	"regexp"
)

const (
	RExAvoid = `[^.()[\]]+`
)

var (
	RExAvoidDo = regexp.MustCompile(RExAvoid)
)

func Parse(caller string) []string {
	return RExAvoidDo.FindAllString(caller, -1)
}
