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

func (m *m) KeyValid() *Key {
	key, err := values.New[*Key](FormatKey(m.KeyComponentsValid()))

	if err != nil {
		errors.Panic(err)
	}

	return key
}

func (m *m) KeyValidWithComponents(components *KeyComponents) *Key {
	key, err := values.New[*Key](FormatKey(components))

	if err != nil {
		errors.Panic(err)
	}

	return key
}

func (m *m) KeyInvalid() {
	_, err := values.New[*Key](FormatKey(m.KeyComponentsInvalid()))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) KeyInvalidWithComponents(components *KeyComponents) {
	_, err := values.New[*Key](FormatKey(components))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) RecipientComponentsValid() *RecipientComponents {
	return &RecipientComponents{
		Service: m.LoremIpsumWord(),
		Entity:  m.LoremIpsumWord(),
		Trigger: m.WordsJoin(m.Words(m.IntRange(1, 2)), "_"),
		Action:  m.LoremIpsumWord(),
		Status:  m.RandomString([]string{"queued", "succeeded", "failed", "done"}),
	}
}

func (m *m) RecipientComponentsInvalid() *RecipientComponents {
	return new(RecipientComponents)
}

func (m *m) RecipientValid() *Recipient {
	recipient, err := values.New[*Recipient](FormatRecipient(m.RecipientComponentsValid()))

	if err != nil {
		errors.Panic(err)
	}

	return recipient
}

func (m *m) RecipientValidWithComponents(components *RecipientComponents) *Recipient {
	recipient, err := values.New[*Recipient](FormatRecipient(components))

	if err != nil {
		errors.Panic(err)
	}

	return recipient
}

func (m *m) RecipientInvalid() {
	_, err := values.New[*Recipient](FormatRecipient(m.RecipientComponentsInvalid()))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) RecipientInvalidWithComponents(components *RecipientComponents) {
	_, err := values.New[*Recipient](FormatRecipient(components))

	if err != nil {
		errors.Panic(err)
	}
}

func (m *m) MessageValid() *Message {
	return &Message{
		ID:         values.Mother().IDValid(),
		OccurredAt: values.Mother().TimeValid(),
		Key:        m.KeyValidWithComponents(m.KeyComponentsValid()),
		Attributes: m.LoremIpsumWord(),
		Meta:       m.LoremIpsumWord(),
	}
}

func (m *m) MessageValidWithKey(key *Key) *Message {
	return &Message{
		ID:         values.Mother().IDValid(),
		OccurredAt: values.Mother().TimeValid(),
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
		ID:         values.Mother().IDValid(),
		OccurredAt: values.Mother().TimeValid(),
		Key:        m.KeyValidWithComponents(m.KeyComponentsValid()),
		Attributes: attributes,
		Meta:       m.LoremIpsumWord(),
	}
}

var Mother = mother.New[m]
