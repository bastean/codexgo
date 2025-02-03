package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var QueryKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Query,
	Entity:  "user",
	Query:   "read",
	Status:  messages.Status.Queued,
})

var ResponseKey = messages.NewKey(&messages.KeyComponents{
	Service:  "user",
	Version:  "1",
	Type:     messages.Type.Response,
	Entity:   "user",
	Response: "read",
	Status:   messages.Status.Done,
})

type QueryAttributes struct {
	ID string
}

type ResponseAttributes struct {
	ID, Email, Username string
	Verified            bool
}

type QueryMeta struct{}

type ResponseMeta struct{}

type Handler struct {
	cases.Read
}

func (h *Handler) Handle(query *messages.Message) (*messages.Message, error) {
	attributes, ok := query.Attributes.(*QueryAttributes)

	if !ok {
		return nil, errors.QueryAssertion("Handle")
	}

	id, err := user.NewID(attributes.ID)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	aggregate, err := h.Read.Run(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := &ResponseAttributes{
		ID:       aggregate.ID.Value,
		Email:    aggregate.Email.Value,
		Username: aggregate.Username.Value,
		Verified: aggregate.Verified.Value,
	}

	return messages.New(
		ResponseKey,
		response,
		new(ResponseMeta),
	), nil
}
