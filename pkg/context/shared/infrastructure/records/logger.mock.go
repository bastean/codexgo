package records

import (
	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func (logger *LoggerMock) Debug(message string) {
	logger.Called(message)
}

func (logger *LoggerMock) Error(message string) {
	logger.Called(message)
}

func (logger *LoggerMock) Fatal(message string) {
	logger.Called(message)
}

func (logger *LoggerMock) Info(message string) {
	logger.Called(message)
}

func (logger *LoggerMock) Success(message string) {
	logger.Called(message)
}
