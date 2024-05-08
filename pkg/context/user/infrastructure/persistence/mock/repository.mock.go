package persistenceMock

import (
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(user *aggregate.User) error {
	args := repository.Called(user)
	return args.Get(0).(error)
}

func (repository *RepositoryMock) Update(user *aggregate.User) error {
	args := repository.Called(user)
	return args.Get(0).(error)
}

func (repository *RepositoryMock) Delete(id sharedModel.ValueObject[string]) error {
	args := repository.Called(id)
	return args.Get(0).(error)
}

func (repository *RepositoryMock) Search(filter model.RepositorySearchCriteria) (*aggregate.User, error) {
	args := repository.Called(filter)
	return args.Get(0).(*aggregate.User), args.Get(1).(error)
}
