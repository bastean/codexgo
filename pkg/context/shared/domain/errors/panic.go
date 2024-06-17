package errors

import (
	"log"
)

func Panic(what, where string) {
	log.Panicf("(%v): [%v]", where, what)
}
