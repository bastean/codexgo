package messages

import (
	"fmt"
	"regexp"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

const (
	RExKeyOrganization = `([a-z0-9]{1,20})`
	RExKeyService      = `([a-z0-9]{1,20})`
	RExKeyVersion      = `(\d+)`
	RExKeyType         = `(event|command|query|response)`
	RExKeyEntity       = `([a-z]{1,20})`
	RExKeyAction       = `([a-z]{1,20})`
	RExKeyStatus       = `(queued|succeeded|failed|done)`
)

var RExKeyComponents = fmt.Sprintf(`^%s\.%s\.%s\.%s\.%s\.%s\.%s$`, RExKeyOrganization, RExKeyService, RExKeyVersion, RExKeyType, RExKeyEntity, RExKeyAction, RExKeyStatus)

var Type = struct {
	Event, Command, Query, Response string
}{
	Event:    "event",
	Command:  "command",
	Query:    "query",
	Response: "response",
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
// Nomenclature of a Key:
//   - (organization).(service).(version).(type).(entity).(action).(status)
//   - codexgo.user.1.event.user.created.succeeded
type KeyComponents struct {
	Organization, Service, Version, Type, Entity, Action, Status string
}

type Key struct {
	values.Object[string]
}

func (k *Key) Validate() error {
	if !regexp.MustCompile(RExKeyComponents).MatchString(k.RawValue()) {
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
