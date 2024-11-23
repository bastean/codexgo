package collection_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
)

type CollectionTestSuite struct {
	suite.Suite
	sut repository.Repository
}

func (s *CollectionTestSuite) SetupTest() {
	session, err := mongodb.Open(
		os.Getenv("CODEXGO_DATABASE_MONGODB_URI"),
		os.Getenv("CODEXGO_DATABASE_MONGODB_NAME"),
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	name := "users-test"

	s.sut, err = collection.Open(session, name)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func (s *CollectionTestSuite) TestCreate() {
	expected := user.RandomPrimitive()

	s.NoError(s.sut.Create(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *CollectionTestSuite) TestCreateErrDuplicateKey() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(aggregate))

	err := s.sut.Create(aggregate)

	var actual *errors.AlreadyExist

	s.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDuplicateKeyError",
		What:  "ID already registered",
		Why: errors.Meta{
			"Field": "ID",
		},
		Who: actual.Who,
	}}

	s.EqualError(expected, actual.Error())
}

func (s *CollectionTestSuite) TestVerify() {
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

func (s *CollectionTestSuite) TestUpdate() {
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

func (s *CollectionTestSuite) TestDelete() {
	aggregate := user.RandomPrimitive()

	s.NoError(s.sut.Create(aggregate))

	s.NoError(s.sut.Delete(aggregate.ID))

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	_, err := s.sut.Search(criteria)

	s.Error(err)
}

func (s *CollectionTestSuite) TestSearch() {
	expected := user.RandomPrimitive()

	s.NoError(s.sut.Create(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *CollectionTestSuite) TestSearchErrDocumentNotFound() {
	aggregate := user.RandomPrimitive()

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	_, err := s.sut.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", aggregate.ID.Value),
		Why: errors.Meta{
			"Index": aggregate.ID.Value,
		},
		Who: actual.Who,
	}}

	s.EqualError(expected, actual.Error())
}

func TestIntegrationCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
