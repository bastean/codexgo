package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
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

func NewRoutingKey(components *RoutingKeyComponents) string {
	if components.Organization == "" {
		components.Organization = "codexgo"
	}

	organization, errOrganization := valueobjs.NewOrganization(components.Organization)
	service, errService := valueobjs.NewService(components.Service)
	version, errVersion := valueobjs.NewVersion(components.Version)
	types, errType := valueobjs.NewType(components.Type)
	entity, errEntity := valueobjs.NewEntity(components.Entity)

	event, errEvent := valueobjs.NewEvent(components.Event)
	command, errCommand := valueobjs.NewCommand(components.Command)

	var action string
	var errAction error

	switch components.Type {
	case Type.Event:
		action = event.Value
		errAction = errEvent
	case Type.Command:
		action = command.Value
		errAction = errCommand
	}

	status, errStatus := valueobjs.NewStatus(components.Status)

	if err := errors.Join(errOrganization, errService, errVersion, errType, errEntity, errAction, errStatus); err != nil {
		errors.Panic(err.Error(), "NewRoutingKey")
	}

	key := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s", organization.Value, service.Value, version.Value, types.Value, entity.Value, action, status.Value)

	key = strings.ToLower(key)

	return key
}
