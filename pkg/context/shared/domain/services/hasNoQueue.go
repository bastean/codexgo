package services

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
)

func HasNoQueue(queues []*queues.Queue, queue *queues.Queue) bool {
	isNotPresent := true

	for _, present := range queues {
		if present.Name == queue.Name {
			isNotPresent = false
			break
		}
	}

	return isNotPresent
}
