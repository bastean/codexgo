package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var QueryKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Query,
	Entity:  "user",
	Query:   "logging",
	Status:  messages.Status.Queued,
})

var ResponseKey = messages.NewKey(&messages.KeyComponents{
	Service:  "user",
	Version:  "1",
	Type:     messages.Type.Response,
	Entity:   "user",
	Response: "logging",
	Status:   messages.Status.Done,
})

type QueryAttributes struct {
	Email, Username, Password string
}

type ResponseAttributes struct {
	ID, Email, Username, Password string
	Verified                      bool
}

type QueryMeta struct{}

type ResponseMeta struct{}

type Handler struct {
	cases.Login
}

func (h *Handler) Handle(query *queries.Query) (*queries.Response, error) {
	attributes, ok := query.Attributes.(*QueryAttributes)

	if !ok {
		return nil, errors.QueryAssertion("Handle")
	}

	var (
		err      error
		email    *user.Email
		username *user.Username
	)

	if attributes.Email != "" {
		email, err = user.NewEmail(attributes.Email)
	}

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	if attributes.Username != "" {
		username, err = user.NewUsername(attributes.Username)
	}

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	plain, err := user.NewPlainPassword(attributes.Password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	found, err := h.Login.Run(email, username, plain)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := (*ResponseAttributes)(found.ToPrimitive())

	return messages.New[queries.Response](
		ResponseKey,
		response,
		new(ResponseMeta),
	), nil
}
