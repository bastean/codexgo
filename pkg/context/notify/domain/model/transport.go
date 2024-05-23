package model

type Transport interface {
	Submit(any) error
}
