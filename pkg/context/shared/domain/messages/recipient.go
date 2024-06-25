package messages

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
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

	var trigger models.ValueObject[string]
	var errTrigger error

	switch {
	case components.Event != "":
		trigger, errTrigger = valueobjs.NewEvent(components.Event)
	case components.Command != "":
		trigger, errTrigger = valueobjs.NewCommand(components.Command)
	}

	status, errStatus := valueobjs.NewStatus(components.Status)

	err := errors.Join(errService, errEntity, errAction, errTrigger, errStatus)

	if err != nil {
		errors.Panic(err.Error(), "NewRecipientName")
	}

	name := fmt.Sprintf("%s.%s.%s_on_%s_%s", service.Value(), entity.Value(), strings.ReplaceAll(action.Value(), " ", "_"), trigger.Value(), status.Value())

	name = strings.ToLower(name)

	return name
}
