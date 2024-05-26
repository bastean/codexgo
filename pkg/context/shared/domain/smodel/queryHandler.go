package smodel

type QueryHandler[Query, Response any] interface {
	Handle(Query) (Response, error)
}
