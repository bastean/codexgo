package messages

import (
	"fmt"
	"regexp"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var RExKey = fmt.Sprintf(`^%s\.%s\.%s\.%s\.%s\.%s\.%s$`, RExOrganization, RExService, RExVersion, RExType, RExEntity, RExAction, RExStatus)

// Terminology:
//   - Organization = Context
//   - Service		= Module
//   - Entity		= Aggregate/Root
//
// Nomenclature of a Key:
//   - (organization).(service).(version).(type).(entity).(action).(status)
//   - codexgo.user.1.event.user.created.succeeded
type KeyComponents struct {
	Organization, Service, Version, Type, Entity, Action, Status string
}

type Key struct {
	values.String
}

func (k *Key) Validate() error {
	if !regexp.MustCompile(RExKey).MatchString(k.RawValue()) {
		errors.Panic(errors.Standard("Key has an invalid nomenclature"))
	}

	k.Valid()

	return nil
}

func ParseKey(key *KeyComponents) string {
	if key.Organization == "" {
		key.Organization = "codexgo"
	}

	return fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s",
		key.Organization,
		key.Service,
		key.Version,
		key.Type,
		key.Entity,
		key.Action,
		key.Status,
	)
}
