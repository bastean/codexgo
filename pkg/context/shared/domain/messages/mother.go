package messages

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) KeyComponentsValid() *KeyComponents {
	return &KeyComponents{
		Service: m.LoremIpsumWord(),
		Version: m.Numerify("#"),
		Type:    m.RandomString([]string{"event", "command", "query", "response"}),
		Entity:  m.LoremIpsumWord(),
		Action:  m.LoremIpsumWord(),
		Status:  m.RandomString([]string{"queued", "succeeded", "failed", "done"}),
	}
}

func (m *m) KeyComponentsInvalid() *KeyComponents {
	return new(KeyComponents)
}

func (m *m) KeyValidWithComponents(components *KeyComponents) *Key {
	key, err := values.New[*Key](ParseKey(components))

	if err != nil {
		errors.Panic(err)
	}

	return key
}

func (m *m) KeyInvalid() {
	_, err := values.New[*Key](ParseKey(m.KeyComponentsInvalid()))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) RecipientComponentsValid() *RecipientComponents {
	return &RecipientComponents{
		Service: m.LoremIpsumWord(),
		Entity:  m.LoremIpsumWord(),
		Trigger: m.WordsJoin(m.Words(m.IntRange(1, 3)), "_"),
		Action:  m.LoremIpsumWord(),
		Status:  m.RandomString([]string{"queued", "succeeded", "failed", "done"}),
	}
}

func (m *m) RecipientComponentsInvalid() *RecipientComponents {
	return new(RecipientComponents)
}

func (m *m) RecipientValidWithComponents(components *RecipientComponents) *Recipient {
	recipient, err := values.New[*Recipient](ParseRecipient(components))

	if err != nil {
		errors.Panic(err)
	}

	return recipient
}

func (m *m) RecipientInvalid() {
	_, err := values.New[*Recipient](ParseRecipient(m.RecipientComponentsInvalid()))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) MessageValid() *Message {
	return &Message{
		ID:         m.UUID(),
		OccurredOn: m.TimeZoneFull(),
		Key:        m.KeyValidWithComponents(m.KeyComponentsValid()),
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
		Key:        m.KeyValidWithComponents(m.KeyComponentsValid()),
		Attributes: attributes,
		Meta:       m.LoremIpsumWord(),
	}
}

var Mother = mother.New[m]()
