package role

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
)

type Transfer interface {
	Submit(*recipient.Recipient) error
}
