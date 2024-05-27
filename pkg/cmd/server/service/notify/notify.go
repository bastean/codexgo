package notify

import (
	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

var SendAccountConfirmation = new(send.Send)

func Init(transport model.Transport) error {
	SendAccountConfirmation.Transport = transport

	return nil
}
