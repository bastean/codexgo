package messages

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Message struct {
	ID         *values.ID
	OccurredAt *values.Time
	Key        *Key
	Attributes any
	Meta       any
}

type Primitive struct {
	ID, OccurredAt, Key *values.StringPrimitive
	Attributes, Meta    any
}

func (m *Message) ToPrimitive() *Primitive {
	return &Primitive{
		ID:         m.ID.ToPrimitive(),
		OccurredAt: m.OccurredAt.ToPrimitive(),
		Key:        m.Key.ToPrimitive(),
		Attributes: m.Attributes,
		Meta:       m.Meta,
	}
}

func FromPrimitive(primitive *Primitive) (*Message, error) {
	id, errID := values.FromPrimitive[*values.ID](primitive.ID)
	occurredAt, errOccurredAt := values.FromPrimitive[*values.Time](primitive.OccurredAt)

	if err := errors.Join(errID, errOccurredAt); err != nil {
		return nil, errors.BubbleUp(err)
	}

	key, _ := values.FromPrimitive[*Key](primitive.Key)

	return &Message{
		ID:         id,
		OccurredAt: occurredAt,
		Key:        key,
		Attributes: primitive.Attributes,
		Meta:       primitive.Meta,
	}, nil
}

func New(key *Key, attributes, meta any) *Message {
	return &Message{
		Key:        key,
		Attributes: attributes,
		Meta:       meta,
	}
}
