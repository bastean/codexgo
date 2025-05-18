package read_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type ReadTestSuite struct {
	suite.Frozen
	SUT        roles.QueryHandler
	read       *read.Case
	repository *persistence.RepositoryMock
}

func (s *ReadTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.read = &read.Case{
		Repository: s.repository,
	}

	s.SUT = &read.Handler{
		Case: s.read,
	}
}

func (s *ReadTestSuite) TestHandle() {
	aggregate := user.Mother.UserValidFromPrimitive()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	response := &read.ResponseAttributes{
		ID:       aggregate.ID.Value(),
		Email:    aggregate.Email.Value(),
		Username: aggregate.Username.Value(),
		Verified: aggregate.Verified.Value(),
	}

	expected := messages.New(
		read.ResponseKey,
		response,
		new(read.ResponseMeta),
	)

	attributes := &read.QueryAttributes{
		ID: aggregate.ID.Value(),
	}

	query := messages.Mother.MessageValidWithAttributes(attributes, false)

	actual, err := s.SUT.Handle(query)

	s.NoError(err)

	s.repository.Mock.AssertExpectations(s.T())

	s.Equal(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
