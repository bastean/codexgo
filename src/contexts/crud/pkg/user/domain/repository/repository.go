package repository

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(email *sharedVO.Email)
	Search(email *sharedVO.Email) (*aggregate.User, error)
}
