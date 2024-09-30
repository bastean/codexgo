package transfers

type Transfer interface {
	Submit(data any) error
}
