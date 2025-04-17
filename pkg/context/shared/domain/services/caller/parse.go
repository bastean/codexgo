package caller

import (
	"regexp"
)

const RExAvoid = `[^.()[\]]+`

func Parse(caller string) []string {
	return regexp.MustCompile(RExAvoid).FindAllString(caller, -1)
}
