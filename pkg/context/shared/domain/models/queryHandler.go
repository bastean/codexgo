package models

type QueryHandler[Query, Response any] interface {
	Handle(Query) (Response, error)
}
