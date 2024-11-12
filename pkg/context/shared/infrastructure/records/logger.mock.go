package records

import (
	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) Debug(message string) {
	m.Called(message)
}

func (m *LoggerMock) Error(message string) {
	m.Called(message)
}

func (m *LoggerMock) Fatal(message string) {
	m.Called(message)
}

func (m *LoggerMock) Info(message string) {
	m.Called(message)
}

func (m *LoggerMock) Success(message string) {
	m.Called(message)
}
