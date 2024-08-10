package services

import (
	"context"
)

type (
	CtxKey string
)

const (
	CtxDefaultKey CtxKey = "default"
)

func SetDefaultContextValue(value any) context.Context {
	return context.WithValue(context.Background(), CtxDefaultKey, value)
}

func GetDefaultContextValue[T any](ctx context.Context) T {
	return ctx.Value(CtxDefaultKey).(T)
}
