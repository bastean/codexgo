package records

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
)

type LoggerMock struct {
	mock.Default
}

func (m *LoggerMock) Debug(format string, values ...any) {
	m.Called(append([]any{format}, values...)...)
}

func (m *LoggerMock) Error(format string, values ...any) {
	m.Called(append([]any{format}, values...)...)
}

func (m *LoggerMock) Fatal(format string, values ...any) {
	m.Called(append([]any{format}, values...)...)
}

func (m *LoggerMock) Info(format string, values ...any) {
	m.Called(append([]any{format}, values...)...)
}

func (m *LoggerMock) Success(format string, values ...any) {
	m.Called(append([]any{format}, values...)...)
}
