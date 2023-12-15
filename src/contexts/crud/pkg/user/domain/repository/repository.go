package repository

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(id *sharedVO.Id)
	Search(id *sharedVO.Id) (*aggregate.User, error)
}
