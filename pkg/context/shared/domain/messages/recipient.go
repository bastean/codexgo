package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
)

// Terminology:
//   - Service = Module
//   - Entity  = Aggregate/Root
//
// Nomenclature of a Recipient Name:
//   - service.entity.action_on_event/command_status
//   - user.user.send_confirmation_on_created_succeeded
type RecipientNameComponents struct {
	Service, Entity, Action, Event, Command, Status string
}

func NewRecipientName(components *RecipientNameComponents) string {
	service, errService := valueobjs.NewService(components.Service)
	entity, errEntity := valueobjs.NewEntity(components.Entity)
	action, errAction := valueobjs.NewAction(components.Action)

	event, errEvent := valueobjs.NewEvent(components.Event)
	command, errCommand := valueobjs.NewCommand(components.Command)

	var trigger string
	var errTrigger error

	switch {
	case components.Event != "":
		trigger = event.Value
		errTrigger = errEvent
	case components.Command != "":
		trigger = command.Value
		errTrigger = errCommand
	}

	status, errStatus := valueobjs.NewStatus(components.Status)

	if err := errors.Join(errService, errEntity, errAction, errTrigger, errStatus); err != nil {
		errors.Panic(err.Error(), "NewRecipientName")
	}

	name := fmt.Sprintf("%s.%s.%s_on_%s_%s", service.Value, entity.Value, strings.ReplaceAll(action.Value, " ", "_"), trigger, status.Value)

	name = strings.ToLower(name)

	return name
}
