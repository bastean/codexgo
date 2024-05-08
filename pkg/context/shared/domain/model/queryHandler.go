package model

type QueryHandler[Query any, Response any] interface {
	Handle(Query) (Response, error)
}
