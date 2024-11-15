package collection_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
)

type CollectionTestSuite struct {
	suite.Suite
	sut     repository.Repository
	hashing *ciphers.HashingMock
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

	s.hashing = new(ciphers.HashingMock)

	s.sut, err = collection.Open(session, name, s.hashing)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func (s *CollectionTestSuite) TestCreate() {
	expected := user.Random()

	expected.Pull()

	s.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	s.NoError(s.sut.Create(expected))

	s.hashing.AssertExpectations(s.T())

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *CollectionTestSuite) TestCreateErrDuplicateKey() {
	account := user.Random()

	s.hashing.On("Hash", account.Password.Value).Return(account.Password.Value)

	s.NoError(s.sut.Create(account))

	err := s.sut.Create(account)

	s.hashing.AssertExpectations(s.T())

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
	account := user.Random()

	s.hashing.On("Hash", account.Password.Value).Return(account.Password.Value)

	s.NoError(s.sut.Create(account))

	s.NoError(s.sut.Verify(account.ID))

	criteria := &repository.SearchCriteria{
		ID: account.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.True(actual.Verified.Value)
}

func (s *CollectionTestSuite) TestUpdate() {
	expected := user.Random()

	expected.Pull()

	s.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	s.NoError(s.sut.Create(expected))

	expected.Password = user.PasswordWithValidValue()

	s.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	s.NoError(s.sut.Update(expected))

	s.hashing.AssertExpectations(s.T())

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *CollectionTestSuite) TestDelete() {
	account := user.Random()

	s.hashing.On("Hash", account.Password.Value).Return(account.Password.Value)

	s.NoError(s.sut.Create(account))

	s.NoError(s.sut.Delete(account.ID))

	criteria := &repository.SearchCriteria{
		ID: account.ID,
	}

	_, err := s.sut.Search(criteria)

	s.Error(err)
}

func (s *CollectionTestSuite) TestSearch() {
	expected := user.Random()

	expected.Pull()

	s.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	s.NoError(s.sut.Create(expected))

	criteria := &repository.SearchCriteria{
		ID: expected.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *CollectionTestSuite) TestSearchErrDocumentNotFound() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		ID: account.ID,
	}

	_, err := s.sut.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", account.ID.Value),
		Why: errors.Meta{
			"Index": account.ID.Value,
		},
		Who: actual.Who,
	}}

	s.EqualError(expected, actual.Error())
}

func TestIntegrationCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
