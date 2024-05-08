package service

import (
	"log"
)

func PanicOnError(where, what string) {
	log.Panicf("(%s): [%s]", where, what)
}
