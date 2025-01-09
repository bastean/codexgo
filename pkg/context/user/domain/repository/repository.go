package repository

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type Criteria struct {
	*user.ID
	*user.Email
	*user.Username
}

type Repository interface {
	Create(*user.User) error
	Update(*user.User) error
	Delete(*user.ID) error
	Search(*Criteria) (*user.User, error)
}
