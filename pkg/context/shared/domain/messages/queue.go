package messages

type (
	BindingKeys = []string
)

type Queue struct {
	Name     string
	Bindings BindingKeys
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
