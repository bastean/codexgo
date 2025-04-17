package errors

import (
	"log"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func Panic(what error) {
	where := "Unknown"

	_, _, name := caller.Received(caller.SkipCurrent)

	if name != "" {
		where = name
	}

	log.Panicf("(%s): %s", where, what)
}
