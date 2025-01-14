package role

type Transfer[T any] interface {
	Submit(data T) error
}
