package persistence

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Create(user *user.User) error {
	m.Called(user)
	return nil
}

func (m *RepositoryMock) Update(user *user.User) error {
	m.Called(user)
	return nil
}

func (m *RepositoryMock) Delete(id *user.ID) error {
	m.Called(id)
	return nil
}

func (m *RepositoryMock) Search(criteria *user.Criteria) (*user.User, error) {
	args := m.Called(criteria)
	return args.Get(0).(*user.User), nil
}
