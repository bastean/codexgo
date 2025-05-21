package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) EventAttributesValid() *EventAttributes {
	return &EventAttributes{
		VerifyToken: recipient.Mother().IDValid().Value(),
		ID:          recipient.Mother().IDValid().Value(),
		Email:       recipient.Mother().EmailValid().Value(),
		Username:    recipient.Mother().UsernameValid().Value(),
	}
}

var Mother = mother.New[m]
