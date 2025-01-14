package role

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type Repository interface {
	Create(*user.User) error
	Update(*user.User) error
	Delete(*user.ID) error
	Search(*user.Criteria) (*user.User, error)
}
