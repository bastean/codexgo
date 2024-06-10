package persistence

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(user *aggregate.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Verify(id models.ValueObject[string]) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Update(user *aggregate.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Delete(id models.ValueObject[string]) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Search(criteria *model.RepositorySearchCriteria) (*aggregate.User, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*aggregate.User), nil
}
