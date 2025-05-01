package messages

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) KeyValid() *Key {
	key, _ := values.New[*Key](ParseKey(&KeyComponents{
		Service: "user",
		Version: "1",
		Type:    Type.Command,
		Entity:  "user",
		Action:  "create",
		Status:  Status.Queued,
	}))

	return key
}

func (m *m) MessageValid() *Message {
	return &Message{
		ID:         m.UUID(),
		OccurredOn: m.TimeZoneFull(),
		Key:        m.KeyValid(),
		Attributes: m.LoremIpsumWord(),
		Meta:       m.LoremIpsumWord(),
	}
}

func (m *m) MessageValidWithKey(key *Key) *Message {
	return &Message{
		ID:         m.UUID(),
		OccurredOn: m.TimeZoneFull(),
		Key:        key,
		Attributes: m.LoremIpsumWord(),
		Meta:       m.LoremIpsumWord(),
	}
}

func (m *m) MessageValidWithAttributes(attributes any, shouldRandomize bool) *Message {
	if shouldRandomize {
		m.StructRandomize(attributes)
	}

	return &Message{
		ID:         m.UUID(),
		OccurredOn: m.TimeZoneFull(),
		Key:        m.KeyValid(),
		Attributes: attributes,
		Meta:       m.LoremIpsumWord(),
	}
}

var Mother = mother.New[m]()
