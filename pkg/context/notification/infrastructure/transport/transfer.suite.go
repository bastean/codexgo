package transport

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
)

type OnlineSuite struct {
	suite.Suite
	SUT       role.Transfer
	Recipient *recipient.Recipient
}

func (s *OnlineSuite) TestSubmit() {
	s.NoError(s.SUT.Submit(s.Recipient))
}

type OfflineSuite struct {
	suite.Suite
	SUT       role.Transfer
	Logger    *records.LoggerMock
	Recipient *recipient.Recipient
	Message   string
}

func (s *OfflineSuite) TestSubmit() {
	s.Logger.Mock.On("Info", s.Message)

	s.NoError(s.SUT.Submit(s.Recipient))

	s.Logger.Mock.AssertExpectations(s.T())
}
