package messages

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func KeyWithValidValue() *Key {
	key, _ := values.New[*Key](ParseKey(&KeyComponents{
		Service: "user",
		Version: "1",
		Type:    Type.Command,
		Entity:  "user",
		Action:  "create",
		Status:  Status.Queued,
	}))

	return key
}
