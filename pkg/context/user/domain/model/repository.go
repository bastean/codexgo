package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type RepositorySearchCriteria struct {
	Id, Email models.ValueObject[string]
}

type Repository interface {
	Save(user *aggregate.User) error
	Verify(id models.ValueObject[string]) error
	Update(user *aggregate.User) error
	Delete(id models.ValueObject[string]) error
	Search(criteria *RepositorySearchCriteria) (*aggregate.User, error)
}
