package repository

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type UserSearchCriteria struct {
	*user.Id
	*user.Email
}

type User interface {
	Save(*user.User) error
	Verify(*user.Id) error
	Update(*user.User) error
	Delete(*user.Id) error
	Search(*UserSearchCriteria) (*user.User, error)
}
