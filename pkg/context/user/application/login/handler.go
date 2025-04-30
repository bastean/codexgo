package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var QueryKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Query,
	Entity:  "user",
	Action:  "login",
	Status:  messages.Status.Queued,
}))

var ResponseKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Response,
	Entity:  "user",
	Action:  "login",
	Status:  messages.Status.Done,
}))

type QueryAttributes struct {
	Email, Username, Password string
}

type ResponseAttributes struct {
	ID, Email, Username string
	Verified            bool
}

type QueryMeta struct{}

type ResponseMeta struct{}

type Handler struct {
	cases.Login
}

func (h *Handler) Handle(query *messages.Message) (*messages.Message, error) {
	attributes, ok := query.Attributes.(*QueryAttributes)

	if !ok {
		return nil, errors.QueryAssertion("Handle")
	}

	if attributes.Email == "" && attributes.Username == "" {
		return nil, errors.New[errors.Failure](&errors.Bubble{
			What: "Email or Username required",
		})
	}

	var (
		err      error
		email    *user.Email
		username *user.Username
	)

	if attributes.Email != "" {
		email, err = values.New[*user.Email](attributes.Email)
	}

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	if attributes.Username != "" {
		username, err = values.New[*user.Username](attributes.Username)
	}

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	plain, err := values.New[*user.PlainPassword](attributes.Password)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate, err := h.Login.Run(email, username, plain)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	response := &ResponseAttributes{
		ID:       aggregate.ID.Value(),
		Email:    aggregate.Email.Value(),
		Username: aggregate.Username.Value(),
		Verified: aggregate.Verified.Value(),
	}

	return messages.New(
		ResponseKey,
		response,
		new(ResponseMeta),
	), nil
}
