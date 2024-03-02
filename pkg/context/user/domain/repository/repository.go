package repository

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
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
