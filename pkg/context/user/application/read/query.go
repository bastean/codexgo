package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const (
	QueryType    queries.Type = "user.query.reading.user"
	ResponseType queries.Type = "user.response.reading.user"
)

type Query struct {
	Id string
}

func (*Query) Type() queries.Type {
	return QueryType
}

type Response struct {
	Id, Email, Username, Password string
	Verified                      bool
}

func (*Response) Type() queries.Type {
	return ResponseType
}

type Handler struct {
	cases.Read
}

func (handler *Handler) SubscribedTo() queries.Type {
	return QueryType
}

func (handler *Handler) ReplyTo() queries.Type {
	return ResponseType
}

func (handler *Handler) Handle(ask queries.Query) (queries.Response, error) {
	data, ok := ask.(*Query)

	if !ok {
		return nil, errors.QueryAssertion("Handle")
	}

	id, err := user.NewId(data.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	found, err := handler.Read.Run(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*found.ToPrimitive())

	return &response, nil
}
