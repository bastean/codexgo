package model

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type RepositorySearchFilter struct {
	Id    *sharedValueObject.Id
	Email *sharedValueObject.Email
}

type Repository interface {
	Save(user *aggregate.User)
	Update(user *aggregate.User)
	Delete(id *sharedValueObject.Id)
	Search(filter RepositorySearchFilter) *aggregate.User
}
