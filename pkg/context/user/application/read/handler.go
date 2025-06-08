package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var QueryKey, _ = values.New[*messages.Key](messages.FormatKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Query,
	Entity:  "user",
	Action:  "read",
	Status:  messages.Status.Queued,
}))

type QueryAttributes = struct {
	ID string
}

type QueryMeta = struct{}

var ResponseKey, _ = values.New[*messages.Key](messages.FormatKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Response,
	Entity:  "user",
	Action:  "read",
	Status:  messages.Status.Done,
}))

type ResponseAttributes = struct {
	ID, Email, Username string
	Verified            bool
}

type ResponseMeta = struct{}

type Handler struct {
	*Case
}

func (h *Handler) Handle(query *messages.Message) (*messages.Message, error) {
	attributes, ok := query.Attributes.(*QueryAttributes)

	if !ok {
		return nil, errors.QueryAssertion()
	}

	user, err := h.Case.Run(attributes)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return messages.New(
		ResponseKey,
		&ResponseAttributes{
			ID:       user.ID.Value(),
			Email:    user.Email.Value(),
			Username: user.Username.Value(),
			Verified: user.Verified.Value(),
		},
		new(ResponseMeta),
	), nil
}
