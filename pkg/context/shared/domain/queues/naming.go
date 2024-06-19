package queues

import (
	"fmt"
	"reflect"
	"strings"
)

type QueueName struct {
	Module, Aggregate, Action, Event, Command string
}

func NewQueueName(queueName *QueueName) string {
	if queueName.Action != "" {
		queueName.Action = strings.ReplaceAll(queueName.Action, " ", "-")
	}

	fields := reflect.ValueOf(*queueName)

	values := []string{}

	for i := 0; i < fields.NumField()-2; i++ {
		value := fields.Field(i).Interface().(string)

		if value != "" {
			values = append(values, value)
		}
	}

	name := strings.Join(values, ".")

	trigger := ""

	switch {
	case queueName.Event != "":
		trigger = strings.ReplaceAll(queueName.Event, " ", ".")
	case queueName.Command != "":
		trigger = strings.ReplaceAll(queueName.Command, " ", ".")
	}

	if trigger != "" {
		name = fmt.Sprintf("%s_on_%s", name, trigger)
	}

	name = strings.ToLower(name)

	return name
}
