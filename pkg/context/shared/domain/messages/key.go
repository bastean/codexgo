package messages

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var (
	RExKey = fmt.Sprintf(`^%s\.%s\.%s\.%s\.%s\.%s\.%s$`, RExOrganization, RExService, RExVersion, RExType, RExEntity, RExAction, RExStatus)
)

var (
	RExKeyDo = regexp.MustCompile(RExKey)
)

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
	if !RExKeyDo.MatchString(k.RawValue()) {
		errors.Panic(errors.Standard("Key has an invalid nomenclature %q", k.RawValue()))
	}

	k.Valid()

	return nil
}

func FormatKey(key *KeyComponents) string {
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

func ParseKey(value string) *KeyComponents {
	_, _ = values.New[*Key](value)

	components := strings.Split(value, ".")

	return &KeyComponents{
		Organization: components[0],
		Service:      components[1],
		Version:      components[2],
		Type:         components[3],
		Entity:       components[4],
		Action:       components[5],
		Status:       components[6],
	}
}
