package log

import (
	"log"
)

type Log struct{}

func (*Log) Debug(message string) {
	log.Println(message)
}

func (*Log) Error(message string) {
	log.Println(message)
}

func (*Log) Fatal(message string) {
	log.Fatal(message)
}

func (*Log) Info(message string) {
	log.Println(message)
}
