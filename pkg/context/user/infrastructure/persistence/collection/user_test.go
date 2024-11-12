package collection_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/collection"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	sut     repository.Repository
	hashing *cryptographic.HashingMock
}

func (s *UserTestSuite) SetupTest() {
	session, err := mongodb.Open(
		os.Getenv("CODEXGO_DATABASE_MONGODB_URI"),
		os.Getenv("CODEXGO_DATABASE_MONGODB_NAME"),
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	name := "users-test"

	s.hashing = new(cryptographic.HashingMock)

	s.sut, err = collection.OpenUser(session, name, s.hashing)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func (s *UserTestSuite) TestCreate() {
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

func (s *UserTestSuite) TestCreateErrDuplicateKey() {
	random := user.Random()

	s.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	s.NoError(s.sut.Create(random))

	err := s.sut.Create(random)

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

func (s *UserTestSuite) TestVerify() {
	random := user.Random()

	s.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	s.NoError(s.sut.Create(random))

	s.NoError(s.sut.Verify(random.ID))

	criteria := &repository.SearchCriteria{
		ID: random.ID,
	}

	actual, err := s.sut.Search(criteria)

	s.NoError(err)

	s.True(actual.Verified.Value)
}

func (s *UserTestSuite) TestUpdate() {
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

func (s *UserTestSuite) TestDelete() {
	random := user.Random()

	s.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	s.NoError(s.sut.Create(random))

	s.NoError(s.sut.Delete(random.ID))

	criteria := &repository.SearchCriteria{
		ID: random.ID,
	}

	_, err := s.sut.Search(criteria)

	s.Error(err)
}

func (s *UserTestSuite) TestSearch() {
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

func (s *UserTestSuite) TestSearchErrDocumentNotFound() {
	random := user.Random()

	criteria := &repository.SearchCriteria{
		ID: random.ID,
	}

	_, err := s.sut.Search(criteria)

	var actual *errors.NotExist

	s.ErrorAs(err, &actual)

	expected := &errors.NotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", random.ID.Value),
		Why: errors.Meta{
			"Index": random.ID.Value,
		},
		Who: actual.Who,
	}}

	s.EqualError(expected, actual.Error())
}

func TestIntegrationUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
