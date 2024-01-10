package cryptographic

import "github.com/stretchr/testify/mock"

type UserHashingMock struct {
	mock.Mock
}

func (m *UserHashingMock) Hash(plain string) string {
	return ""
}

func (m *UserHashingMock) IsNotEqual(hashed, plain string) bool {
	return false
}
