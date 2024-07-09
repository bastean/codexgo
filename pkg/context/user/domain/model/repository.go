package model

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
)

type RepositorySearchCriteria struct {
	*user.Id
	*user.Email
}

type Repository interface {
	Save(*user.User) error
	Verify(*user.Id) error
	Update(*user.User) error
	Delete(*user.Id) error
	Search(*RepositorySearchCriteria) (*user.User, error)
}
