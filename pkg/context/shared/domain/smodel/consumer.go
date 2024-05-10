package smodel

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
)

type Consumer interface {
	SubscribedTo() []*squeue.Queue
	On(message *smessage.Message) error
}
