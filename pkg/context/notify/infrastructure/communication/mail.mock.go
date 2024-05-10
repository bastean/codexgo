package communication

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/stretchr/testify/mock"
)

type MailMock struct {
	mock.Mock
}

func (mail *MailMock) Send(template model.MailTemplate) error {
	mail.Called(template)
	return nil
}
