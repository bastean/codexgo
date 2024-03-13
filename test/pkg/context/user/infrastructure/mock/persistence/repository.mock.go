package persistenceMock

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
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

func (m *RepositoryMock) Search(filter model.RepositorySearchFilter) *aggregate.User {
	args := m.Called(filter)
	return args.Get(0).(*aggregate.User)
}

func NewRepositoryMock() *RepositoryMock {
	return new(RepositoryMock)
}
