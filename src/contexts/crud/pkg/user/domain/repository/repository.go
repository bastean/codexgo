package repository

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type Filter struct {
	Id    *sharedVO.Id
	Email *sharedVO.Email
}

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(id *sharedVO.Id)
	Search(filter Filter) *aggregate.User
}
