package messages

type Queue struct {
	Name     string
	Bindings []string
}

func HasNoQueue(queues []Queue, queue *Queue) bool {
	isNotPresent := true

	for _, present := range queues {
		if present.Name == queue.Name {
			isNotPresent = false
			break
		}
	}

	return isNotPresent
}
