package communicationMock

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/stretchr/testify/mock"
)

type MailMock struct {
	mock.Mock
}

func (mail *MailMock) Send(template model.MailTemplate) error {
	args := mail.Called(template)
	return args.Get(0).(error)
}
