package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
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

	var action models.ValueObject[string]
	var errAction error

	switch components.Type {
	case Type.Event:
		action, errAction = valueobjs.NewEvent(components.Event)
	case Type.Command:
		action, errAction = valueobjs.NewCommand(components.Command)
	}

	status, errStatus := valueobjs.NewStatus(components.Status)

	err := errors.Join(errOrganization, errService, errVersion, errType, errEntity, errAction, errStatus)

	if err != nil {
		errors.Panic(err.Error(), "NewRoutingKey")
	}

	key := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s", organization.Value(), service.Value(), version.Value(), types.Value(), entity.Value(), action.Value(), status.Value())

	key = strings.ToLower(key)

	return key
}
