package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
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
// Nomenclature of a Routing Key (Topic):
//   - organization.service.version.type.entity.event/command.status
//   - codexgo.user.1.event.user.created.succeeded
type RoutingKeyComponents struct {
	Organization, Service, Version, Type, Entity, Event, Command, Status string
}

func NewRoutingKey(routing *RoutingKeyComponents) string {
	if routing.Organization == "" {
		routing.Organization = "codexgo"
	}

	organization, errOrganization := components.NewOrganization(routing.Organization)
	service, errService := components.NewService(routing.Service)
	version, errVersion := components.NewVersion(routing.Version)
	types, errType := components.NewType(routing.Type)
	entity, errEntity := components.NewEntity(routing.Entity)

	event, errEvent := components.NewEvent(routing.Event)
	command, errCommand := components.NewCommand(routing.Command)

	var action string
	var errAction error

	switch routing.Type {
	case Type.Event:
		action = event.Value
		errAction = errEvent
	case Type.Command:
		action = command.Value
		errAction = errCommand
	}

	status, errStatus := components.NewStatus(routing.Status)

	if err := errors.Join(errOrganization, errService, errVersion, errType, errEntity, errAction, errStatus); err != nil {
		errors.Panic(err.Error(), "NewRoutingKey")
	}

	key := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s", organization.Value, service.Value, version.Value, types.Value, entity.Value, action, status.Value)

	key = strings.ToLower(key)

	return key
}
