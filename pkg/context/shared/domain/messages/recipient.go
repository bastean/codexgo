package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
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

func NewRecipientName(recipient *RecipientNameComponents) string {
	service, errService := components.NewService(recipient.Service)
	entity, errEntity := components.NewEntity(recipient.Entity)
	action, errAction := components.NewAction(recipient.Action)

	event, errEvent := components.NewEvent(recipient.Event)
	command, errCommand := components.NewCommand(recipient.Command)

	var trigger string
	var errTrigger error

	switch {
	case recipient.Event != "":
		trigger = event.Value
		errTrigger = errEvent
	case recipient.Command != "":
		trigger = command.Value
		errTrigger = errCommand
	}

	status, errStatus := components.NewStatus(recipient.Status)

	if err := errors.Join(errService, errEntity, errAction, errTrigger, errStatus); err != nil {
		errors.Panic(err.Error(), "NewRecipientName")
	}

	name := fmt.Sprintf("%s.%s.%s_on_%s_%s", service.Value, entity.Value, strings.ReplaceAll(action.Value, " ", "_"), trigger, status.Value)

	name = strings.ToLower(name)

	return name
}
