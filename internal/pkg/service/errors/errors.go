package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	Bubble = errors.Bubble
)

type (
	ErrInvalidValue = errors.ErrInvalidValue
	ErrAlreadyExist = errors.ErrAlreadyExist
	ErrNotExist     = errors.ErrNotExist
	ErrFailure      = errors.ErrFailure
	ErrInternal     = errors.ErrInternal
)

var (
	NewFailure  = errors.NewFailure
	NewInternal = errors.NewInternal
)

var (
	Panic    = errors.Panic
	BubbleUp = errors.BubbleUp
)

var (
	As = errors.As
	Is = errors.Is
)
