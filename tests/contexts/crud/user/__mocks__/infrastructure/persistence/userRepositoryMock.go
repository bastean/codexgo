package persistence

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Save(user *aggregate.User) {
	m.Called(user)
}

func (m *UserRepositoryMock) Update(user *aggregate.User) {
	m.Called(user)
}

func (m *UserRepositoryMock) Delete(id *sharedVO.Id) {
	m.Called(id)
}

func (m *UserRepositoryMock) Search(filter repository.Filter) *aggregate.User {
	args := m.Called(filter)
	return args.Get(0).(*aggregate.User)
}
