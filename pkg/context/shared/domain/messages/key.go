package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

var Type = struct {
	Event, Command string
}{
	Event:   "event",
	Command: "command",
}

var Status = struct {
	Queued, Succeeded, Failed, Done string
}{
	Queued:    "queued",
	Succeeded: "succeeded",
	Failed:    "failed",
	Done:      "done",
}

// Terminology:
//   - Organization = Context
//   - Service		= Module
//   - Entity		= Aggregate/Root
//
// Nomenclature of a Key:
//   - organization.service.version.type.entity.event/command.status
//   - codexgo.user.1.event.user.created.succeeded
type KeyComponents struct {
	Organization, Service, Version, Type, Entity, Event, Command, Status string
}

func NewKey(key *KeyComponents) events.Key {
	if key.Organization == "" {
		key.Organization = "codexgo"
	}

	organization, errOrganization := components.NewOrganization(key.Organization)
	service, errService := components.NewService(key.Service)
	version, errVersion := components.NewVersion(key.Version)
	types, errType := components.NewType(key.Type)
	entity, errEntity := components.NewEntity(key.Entity)

	event, errEvent := components.NewEvent(key.Event)
	command, errCommand := components.NewCommand(key.Command)

	var action string
	var errAction error

	switch key.Type {
	case Type.Event:
		action = event.Value
		errAction = errEvent
	case Type.Command:
		action = command.Value
		errAction = errCommand
	}

	status, errStatus := components.NewStatus(key.Status)

	if err := errors.Join(errOrganization, errService, errVersion, errType, errEntity, errAction, errStatus); err != nil {
		errors.Panic(err.Error(), "NewRoutingKey")
	}

	value := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s", organization.Value, service.Value, version.Value, types.Value, entity.Value, action, status.Value)

	value = strings.ToLower(value)

	return events.Key(value)
}
