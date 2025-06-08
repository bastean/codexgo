package persistence

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type RepositorySuite struct {
	suite.Default
	SUT role.Repository
}

func (s *RepositorySuite) TestCreate() {
	expected := user.Mother().UserValid()

	s.NoError(s.SUT.Create(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	expected.Pull()

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestCreateErrDuplicateValue() {
	registered := user.Mother().UserValid()

	aggregate := user.Mother().UserValid()

	s.NoError(s.SUT.Create(registered))

	aggregate.ID = registered.ID

	err := s.SUT.Create(aggregate)

	var actual *errors.AlreadyExist

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "HandleErrDuplicateValue")

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: actual.Where,
		What:  "ID already registered",
		Why: errors.Meta{
			"Field": "ID",
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestUpdate() {
	var err error

	expected := user.Mother().UserValid()

	s.NoError(s.SUT.Create(expected))

	expected.Password, err = values.Replace(expected.Password, user.Mother().PasswordValid().Value())

	s.NoError(err)

	s.NoError(s.SUT.Update(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	expected.Pull()

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestUpdateErrDuplicateValue() {
	registered := user.Mother().UserValid()

	aggregate := user.Mother().UserValid()

	s.NoError(s.SUT.Create(registered))

	s.NoError(s.SUT.Create(aggregate))

	aggregate.Email = registered.Email

	err := s.SUT.Update(aggregate)

	var actual *errors.AlreadyExist

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "HandleErrDuplicateValue")

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: actual.Where,
		What:  "Email already registered",
		Why: errors.Meta{
			"Field": "Email",
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestDelete() {
	aggregate := user.Mother().UserValid()

	s.NoError(s.SUT.Create(aggregate))

	s.NoError(s.SUT.Delete(aggregate.ID))

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "HandleErrNotFound")

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: actual.Where,
		What:  aggregate.ID.Value() + " not found",
		Why: errors.Meta{
			"Index": aggregate.ID.Value(),
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearch() {
	expected := user.Mother().UserValid()

	s.NoError(s.SUT.Create(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	expected.Pull()

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearchErrCriteria() {
	criteria := new(user.Criteria)

	_, err := s.SUT.Search(criteria)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "Search")

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: actual.Where,
		What:  "Criteria not defined",
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearchErrNotFound() {
	aggregate := user.Mother().UserValid()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "HandleErrNotFound")

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: actual.Where,
		What:  aggregate.ID.Value() + " not found",
		Why: errors.Meta{
			"Index": aggregate.ID.Value(),
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}
