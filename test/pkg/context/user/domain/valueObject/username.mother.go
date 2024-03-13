package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	"github.com/bastean/codexgo/test/pkg/context/shared/domain/service"
)

func NewUsername(username string) *valueObject.Username {
	return valueObject.NewUsername(username)
}

func RandomUsername() *valueObject.Username {
	return NewUsername(service.Mother.Regex(`[a-z0-9]{2,20}`))
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
