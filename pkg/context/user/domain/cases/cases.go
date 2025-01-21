package cases

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type (
	Create interface {
		Run(*user.User) error
	}
	Read interface {
		Run(*user.ID) (*user.User, error)
	}
	Update interface {
		Run(*user.User, *user.PlainPassword) error
	}
	Delete interface {
		Run(*user.ID, *user.PlainPassword) error
	}
	Verify interface {
		Run(*user.ID, *user.ID) error
	}
	Forgot interface {
		Run(*user.ID, *user.Email) (*user.User, error)
	}
	Reset interface {
		Run(*user.ID, *user.ID, *user.PlainPassword) error
	}
	Login interface {
		Run(*user.Email, *user.Username, *user.PlainPassword) (*user.User, error)
	}
)
