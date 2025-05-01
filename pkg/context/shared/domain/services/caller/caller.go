package caller

import (
	"path/filepath"
	"runtime"
)

const (
	FromCurrent = 0
	SkipCurrent = 1
)

func Received(skip int) (pkg, receiver, method string) {
	pc, _, _, _ := runtime.Caller(skip + 1)

	if caller := runtime.FuncForPC(pc); caller != nil {
		names := Parse(filepath.Base(caller.Name()))

		switch len(names) {
		case 2:
			pkg = names[0]
			method = names[1]
		case 3:
			pkg = names[0]
			receiver = names[1]
			method = names[2]
		}
	}

	return pkg, receiver, method
}
