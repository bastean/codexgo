package table_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/sqlite"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/sqlite/table"
)

type TableTestSuite struct {
	suite.Suite
	sut repository.Repository
}

func (s *TableTestSuite) SetupTest() {
	session, err := sqlite.Open(os.Getenv("CODEXGO_DATABASE_SQLITE_DSN"))

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	s.sut, err = table.Open(session)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func (*TableTestSuite) TearDownTest() {
	os.Remove(os.Getenv("CODEXGO_DATABASE_SQLITE_DSN"))
}

func (s *TableTestSuite) TestCreate() {
	expected := user.RandomPrimitive()

	s.NoError(s.sut.Create(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *TableTestSuite) TestCreateErrDuplicateValue() {
	registered := user.RandomPrimitive()

	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(registered))

	aggregate.ID = registered.ID

	err := s.sut.Create(aggregate)

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

	s.EqualError(expected, actual.Error())
}

func (s *TableTestSuite) TestVerify() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(aggregate))

	s.NoError(s.sut.Verify(aggregate.ID))

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.True(actual.Verified.Value)
}

func (s *TableTestSuite) TestUpdate() {
	expected := user.RandomPrimitive()

	s.NoError(s.sut.Create(expected))

	expected.CipherPassword = user.CipherPasswordWithValidValue()

	s.NoError(s.sut.Update(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *TableTestSuite) TestUpdateErrDuplicateValue() {
	registered := user.RandomPrimitive()

	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(registered))

	s.NoError(s.sut.Create(aggregate))

	aggregate.Email = registered.Email

	err := s.sut.Update(aggregate)

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

	s.EqualError(expected, actual.Error())
}

func (s *TableTestSuite) TestDelete() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(aggregate))

	s.NoError(s.sut.Delete(aggregate.ID))

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	_, err := s.sut.Search(criteria)

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

	s.EqualError(expected, actual.Error())
}

func (s *TableTestSuite) TestSearch() {
	expected := user.RandomPrimitive()

	s.NoError(s.sut.Create(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *TableTestSuite) TestSearchErrNotFound() {
	aggregate := user.RandomPrimitive()

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	_, err := s.sut.Search(criteria)

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

	s.EqualError(expected, actual.Error())
}

func TestIntegrationTableSuite(t *testing.T) {
	suite.Run(t, new(TableTestSuite))
}
