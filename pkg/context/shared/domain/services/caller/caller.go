package caller

import (
	"path/filepath"
	"runtime"
	"strings"
)

const (
	FromCurrent = 0
	SkipCurrent = 1
)

const (
	Separator = "/"
)

const (
	DefaultWhere    = "UNKNOWN"
	DefaultPkg      = "UNKNOWN"
	DefaultReceiver = "UNKNOWN"
	DefaultMethod   = "UNKNOWN"
)

func Received(skip int) (where, pkg, receiver, method string) {
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

	where = strings.Join(strings.Fields(strings.Join([]string{pkg, receiver, method}, " ")), Separator)

	if where == "" {
		where = DefaultWhere
	}

	if pkg == "" {
		pkg = DefaultPkg
	}

	if receiver == "" {
		receiver = DefaultReceiver
	}

	if method == "" {
		method = DefaultMethod
	}

	return where, pkg, receiver, method
}
