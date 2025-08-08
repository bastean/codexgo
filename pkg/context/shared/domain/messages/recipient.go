package messages

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var (
	RExRecipient = fmt.Sprintf(`^%s\.%s\.%s_on_%s_%s$`, RExService, RExEntity, RExTrigger, RExAction, RExStatus)
)

var (
	RExRecipientDo = regexp.MustCompile(RExRecipient)
)

// Terminology:
//   - Service = Module
//   - Entity  = Aggregate/Root
//
// Nomenclature of a Recipient:
//   - (service).(entity).(trigger)_on_(action)_(status)
//   - user.user.send_confirmation_on_created_succeeded
type RecipientComponents struct {
	Service, Entity, Trigger, Action, Status string
}

type Recipient struct {
	values.String
}

func (r *Recipient) Validate() error {
	if !RExRecipientDo.MatchString(r.RawValue()) {
		errors.Panic(errors.Standard("Recipient has an invalid nomenclature %q", r.RawValue()))
	}

	r.Valid()

	return nil
}

func FormatRecipient(recipient *RecipientComponents) string {
	return fmt.Sprintf("%s.%s.%s_on_%s_%s",
		recipient.Service,
		recipient.Entity,
		recipient.Trigger,
		recipient.Action,
		recipient.Status,
	)
}

func ParseRecipient(value string) *RecipientComponents {
	_, _ = values.New[*Recipient](value)

	components := strings.Split(value, ".")

	recipient := &RecipientComponents{
		Service: components[0],
		Entity:  components[1],
	}

	underscores := strings.Split(components[2], "_on_")

	recipient.Trigger = underscores[0]

	underscores = strings.Split(underscores[1], "_")

	recipient.Action = underscores[0]

	recipient.Status = underscores[1]

	return recipient
}
