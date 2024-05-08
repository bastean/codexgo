package model

type ValueObject[Value any] interface {
	Value() Value
	IsValid() error
}
