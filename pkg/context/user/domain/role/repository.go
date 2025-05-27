package role

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type Repository interface {
	Create(*user.User) error
	Update(*user.User) error
	Delete(*values.ID) error
	Search(*user.Criteria) (*user.User, error)
}
