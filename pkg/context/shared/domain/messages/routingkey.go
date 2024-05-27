package messages

import (
	"reflect"
	"strings"
)

var Type = struct {
	Event   string
	Command string
}{
	Event:   "event",
	Command: "command",
}

var Status = struct {
	Queued    string
	Succeeded string
	Failed    string
	Done      string
}{
	Queued:    "queued",
	Succeeded: "succeeded",
	Failed:    "failed",
	Done:      "done",
}

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

func NewRoutingKey(routingKey *MessageRoutingKey) string {
	if routingKey.Organization == "" {
		routingKey.Organization = "bastean"
	}

	switch routingKey.Type {
	case Type.Event:
		routingKey.Command = ""
	case Type.Command:
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
