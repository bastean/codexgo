package messages

import (
	"fmt"
	"regexp"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var RExRecipient = fmt.Sprintf(`^%s\.%s\.%s_on_%s_%s$`, RExService, RExEntity, RExTrigger, RExAction, RExStatus)

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
	if !regexp.MustCompile(RExRecipient).MatchString(r.RawValue()) {
		errors.Panic(errors.Standard("Recipient has an invalid nomenclature %q", r.RawValue()))
	}

	r.Valid()

	return nil
}

func ParseRecipient(recipient *RecipientComponents) string {
	return fmt.Sprintf("%s.%s.%s_on_%s_%s",
		recipient.Service,
		recipient.Entity,
		recipient.Trigger,
		recipient.Action,
		recipient.Status,
	)
}
