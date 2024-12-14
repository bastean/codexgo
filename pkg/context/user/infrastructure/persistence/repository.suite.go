package persistence

import (
	"fmt"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type RepositorySuite struct {
	suite.Suite
	SUT repository.Repository
}

func (s *RepositorySuite) TestCreate() {
	expected := user.RandomPrimitive()

	s.NoError(s.SUT.Create(expected))

	criteria := &repository.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestCreateErrDuplicateValue() {
	registered := user.RandomPrimitive()

	aggregate := user.RandomPrimitive()

	s.NoError(s.SUT.Create(registered))

	aggregate.ID = registered.ID

	err := s.SUT.Create(aggregate)

	var actual *errors.AlreadyExist

	s.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrDuplicateValue",
		What:  "ID already registered",
		Why: errors.Meta{
			"Field": "ID",
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestVerify() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.SUT.Create(aggregate))

	s.NoError(s.SUT.Verify(aggregate.ID))

	criteria := &repository.Criteria{
		ID: aggregate.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.True(actual.Verified.Value)
}

func (s *RepositorySuite) TestUpdate() {
	expected := user.RandomPrimitive()

	s.NoError(s.SUT.Create(expected))

	expected.CipherPassword = user.CipherPasswordWithValidValue()

	s.NoError(s.SUT.Update(expected))

	criteria := &repository.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestUpdateErrDuplicateValue() {
	registered := user.RandomPrimitive()

	aggregate := user.RandomPrimitive()

	s.NoError(s.SUT.Create(registered))

	s.NoError(s.SUT.Create(aggregate))

	aggregate.Email = registered.Email

	err := s.SUT.Update(aggregate)

	var actual *errors.AlreadyExist

	s.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrDuplicateValue",
		What:  "Email already registered",
		Why: errors.Meta{
			"Field": "Email",
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestDelete() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.SUT.Create(aggregate))

	s.NoError(s.SUT.Delete(aggregate.ID))

	criteria := &repository.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrNotFound",
		What:  fmt.Sprintf("%s not found", aggregate.ID.Value),
		Why: errors.Meta{
			"Index": aggregate.ID.Value,
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearch() {
	expected := user.RandomPrimitive()

	s.NoError(s.SUT.Create(expected))

	criteria := &repository.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearchErrCriteria() {
	criteria := new(repository.Criteria)

	_, err := s.SUT.Search(criteria)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Search",
		What:  "Criteria not defined",
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearchErrNotFound() {
	aggregate := user.RandomPrimitive()

	criteria := &repository.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrNotFound",
		What:  fmt.Sprintf("%s not found", aggregate.ID.Value),
		Why: errors.Meta{
			"Index": aggregate.ID.Value,
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}