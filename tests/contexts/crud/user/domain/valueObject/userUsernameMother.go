package valueObject

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewUsername(username string) *valueObject.Username {
	return valueObject.NewUsername(username)
}

func RandomUsername() *valueObject.Username {
	return NewUsername(mother.Creator.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() *valueObject.Username {
	return NewUsername("x")
}

func WithInvalidUsernameAlphanumeric() *valueObject.Username {
	return NewUsername("<></>")
}

func EmptyUsername() *valueObject.Username {
	return NewUsername("")
}
