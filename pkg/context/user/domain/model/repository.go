package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type RepositorySearchCriteria struct {
	Id    model.ValueObject[string]
	Email model.ValueObject[string]
}

type Repository interface {
	Save(user *aggregate.User) error
	Update(user *aggregate.User) error
	Delete(id model.ValueObject[string]) error
	Search(filter RepositorySearchCriteria) (*aggregate.User, error)
}
