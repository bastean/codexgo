package cryptographic

import (
	"github.com/stretchr/testify/mock"
)

type HashingMock struct {
	mock.Mock
}

func (hashing *HashingMock) Hash(plain string) (string, error) {
	args := hashing.Called(plain)
	return args.Get(0).(string), nil
}

func (hashing *HashingMock) IsNotEqual(hashed, plain string) bool {
	args := hashing.Called(hashed, plain)
	return args.Get(0).(bool)
}
