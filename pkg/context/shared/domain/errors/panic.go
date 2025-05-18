package errors

import (
	"log"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func Panic(what error) {
	_, _, where := caller.Received(caller.SkipCurrent)
	log.Panicf("(%s): %s", where, what)
}
