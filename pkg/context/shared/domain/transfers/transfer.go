package transfers

type Transfer interface {
	Submit(any) error
}
