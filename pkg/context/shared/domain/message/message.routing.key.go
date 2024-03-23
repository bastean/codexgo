package message

import (
	"reflect"
	"strings"
)

var (
	Event   = "event"
	Command = "command"
)

var (
	Queued    = "queued"
	Succeeded = "succeeded"
	Failed    = "failed"
	Done      = "done"
)

type MessageRoutingKey struct {
	Organization string
	Module       string
	Version      string
	Type         string
	Aggregate    string
	Event        string
	Command      string
	Status       string
}

func NewMessageRoutingKey(routingKey *MessageRoutingKey) string {
	if routingKey.Organization == "" {
		routingKey.Organization = "bastean"
	}

	switch routingKey.Type {
	case Event:
		routingKey.Command = ""
	case Command:
		routingKey.Event = ""
	}

	fields := reflect.ValueOf(*routingKey)

	values := []string{}

	for i := 0; i < fields.NumField(); i++ {
		value := fields.Field(i).Interface().(string)

		if value != "" {
			values = append(values, value)
		}
	}

	key := strings.Join(values, ".")

	key = strings.ToLower(key)

	return key
}
