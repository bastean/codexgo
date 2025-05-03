package persistence

import (
	"os"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type RepositorySuite struct {
	suite.Suite
	SUT role.Repository
}

func (s *RepositorySuite) SetupTest() {
	s.NoError(os.Setenv("GOTEST_FROZEN", "1"))
}

func (s *RepositorySuite) TestCreate() {
	expected := user.Mother.UserValid()

	s.NoError(s.SUT.Create(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestCreateErrDuplicateValue() {
	registered := user.Mother.UserValid()

	aggregate := user.Mother.UserValid()

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

func (s *RepositorySuite) TestUpdate() {
	expected := user.Mother.UserValid()

	s.NoError(s.SUT.Create(expected))

	expected.CipherPassword = user.Mother.CipherPasswordValid()

	s.NoError(s.SUT.Update(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	actual.Updated = expected.Updated

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestUpdateErrDuplicateValue() {
	registered := user.Mother.UserValid()

	aggregate := user.Mother.UserValid()

	s.NoError(s.SUT.Create(registered))

	s.NoError(s.SUT.Create(aggregate))

	aggregate.Email = registered.Email

	s.NoError(os.Unsetenv("GOTEST_FROZEN"))

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
	aggregate := user.Mother.UserValid()

	s.NoError(s.SUT.Create(aggregate))

	s.NoError(s.SUT.Delete(aggregate.ID))

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrNotFound",
		What:  aggregate.ID.Value() + " not found",
		Why: errors.Meta{
			"Index": aggregate.ID.Value(),
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearch() {
	expected := user.Mother.UserValid()

	s.NoError(s.SUT.Create(expected))

	criteria := &user.Criteria{
		ID: expected.ID,
	}

	actual, err := s.SUT.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TestSearchErrCriteria() {
	criteria := new(user.Criteria)

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
	aggregate := user.Mother.UserValid()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	_, err := s.SUT.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleErrNotFound",
		What:  aggregate.ID.Value() + " not found",
		Why: errors.Meta{
			"Index": aggregate.ID.Value(),
		},
		Who: actual.Who,
	}}

	s.Equal(expected, actual)
}

func (s *RepositorySuite) TearDownTest() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN"))
}
