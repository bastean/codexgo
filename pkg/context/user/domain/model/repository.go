package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type RepositorySearchCriteria struct {
	Id    smodel.ValueObject[string]
	Email smodel.ValueObject[string]
}

type Repository interface {
	Save(user *aggregate.User) error
	Update(user *aggregate.User) error
	Delete(id smodel.ValueObject[string]) error
	Search(filter RepositorySearchCriteria) (*aggregate.User, error)
}
