package sservice

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
)

func HasNoQueue(queues []*squeue.Queue, queue *squeue.Queue) bool {
	isNotPresent := true

	for _, present := range queues {
		if present.Name == queue.Name {
			isNotPresent = false
			break
		}
	}

	return isNotPresent
}
