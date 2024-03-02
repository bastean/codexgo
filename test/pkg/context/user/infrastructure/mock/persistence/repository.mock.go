package persistence

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Save(user *aggregate.User) {
	m.Called(user)
}

func (m *RepositoryMock) Update(user *aggregate.User) {
	m.Called(user)
}

func (m *RepositoryMock) Delete(id *sharedVO.Id) {
	m.Called(id)
}

func (m *RepositoryMock) Search(filter repository.Filter) *aggregate.User {
	args := m.Called(filter)
	return args.Get(0).(*aggregate.User)
}
