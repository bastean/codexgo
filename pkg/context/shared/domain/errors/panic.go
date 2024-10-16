package errors

import (
	"log"
)

func Panic(what, where string) {
	log.Panicf("(%s): %s", where, what)
}
