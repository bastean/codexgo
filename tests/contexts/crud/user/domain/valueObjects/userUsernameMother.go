package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewUsername(username string) *valueObjects.Username {
	return valueObjects.NewUsername(username)
}

func RandomUsername() *valueObjects.Username {
	return NewUsername(mother.Creator.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() *valueObjects.Username {
	return NewUsername("x")
}

func WithInvalidUsernameAlphanumeric() *valueObjects.Username {
	return NewUsername("<></>")
}

func EmptyUsername() *valueObjects.Username {
	return NewUsername("")
}
