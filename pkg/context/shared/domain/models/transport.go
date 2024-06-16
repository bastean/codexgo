package models

type Transport interface {
	Submit(any) error
}
