package records

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
)

type LoggerMock struct {
	mock.Default
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
