package ascii

import (
	"strings"
)

func FixWidth(values ...[]string) {
	for _, lines := range values {
		var width, maxWidth int

		for _, line := range lines {
			width = len(line)

			if width > maxWidth {
				maxWidth = width
			}
		}

		for i, line := range lines {
			width = len(line)

			if width < maxWidth {
				lines[i] += strings.Repeat(" ", (maxWidth - width))
			}
		}
	}
}
