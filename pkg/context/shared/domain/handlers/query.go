package handlers

type Query[Query, Response any] interface {
	Handle(Query) (Response, error)
}
