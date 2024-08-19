package persistence

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (repository *UserMock) Create(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Verify(id *user.Id) error {
	repository.Called(id)
	return nil
}

func (repository *UserMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Delete(id *user.Id) error {
	repository.Called(id)
	return nil
}

func (repository *UserMock) Search(criteria *repository.SearchCriteria) (*user.User, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*user.User), nil
}
