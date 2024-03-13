package model

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type RepositorySearchFilter struct {
	Id    *sharedVO.Id
	Email *sharedVO.Email
}

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(id *sharedVO.Id)
	Search(filter RepositorySearchFilter) *aggregate.User
}
