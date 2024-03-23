package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomUsername() *valueObject.Username {
	return valueObject.NewUsername(mother.Create.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() *valueObject.Username {
	return valueObject.NewUsername("x")
}

func WithInvalidUsernameAlphanumeric() *valueObject.Username {
	return valueObject.NewUsername("<></>")
}

func EmptyUsername() *valueObject.Username {
	return valueObject.NewUsername("")
}
