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
	Email, Password string
}

type ResponseAttributes struct {
	Id, Email, Username, Password string
	Verified                      bool
}

type QueryMeta struct{}

type ResponseMeta struct{}

type Handler struct {
	cases.Login
}

func (handler *Handler) Handle(query *queries.Query) (*queries.Response, error) {
	attributes, ok := query.Attributes.(*QueryAttributes)

	if !ok {
		return nil, errors.QueryAssertion("Handle")
	}

	email, errEmail := user.NewEmail(attributes.Email)

	password, errPassword := user.NewPassword(attributes.Password)

	err := errors.Join(errEmail, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	found, err := handler.Login.Run(email, password)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := ResponseAttributes(*found.ToPrimitive())

	return messages.New[queries.Response](
		ResponseKey,
		&response,
		new(ResponseMeta),
	), nil
}
