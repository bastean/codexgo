package transport

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type OnlineSuite[T any] struct {
	suite.Suite
	SUT        role.Transfer[T]
	Attributes T
}

func (s *OnlineSuite[T]) TestSubmit() {
	s.NoError(s.SUT.Submit(s.Attributes))
}

type OfflineSuite[T any] struct {
	suite.Suite
	SUT        role.Transfer[T]
	Logger     *records.LoggerMock
	Attributes T
	Message    string
}

func (s *OfflineSuite[T]) TestSubmit() {
	s.Logger.Mock.On("Info", s.Message)

	s.NoError(s.SUT.Submit(s.Attributes))

	s.Logger.Mock.AssertExpectations(s.T())
}
