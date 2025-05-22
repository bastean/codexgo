package mock

import (
	"github.com/stretchr/testify/mock"
)

type (
	Arguments = mock.Arguments
)

type Default struct {
	mock.Mock
}
