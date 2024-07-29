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
	sut     repository.User
	hashing *cryptographic.HashingMock
}

func (suite *UserTestSuite) SetupTest() {
	session, _ := mongodb.Open(
		os.Getenv("DATABASE_MONGODB_URI"),
		os.Getenv("DATABASE_MONGODB_NAME"),
	)

	name := "users-test"

	suite.hashing = new(cryptographic.HashingMock)

	suite.sut, _ = collection.OpenUser(session, name, suite.hashing)
}

func (suite *UserTestSuite) TestSave() {
	expected := user.Random()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NoError(suite.sut.Save(expected))

	suite.hashing.AssertExpectations(suite.T())

	criteria := &repository.UserSearchCriteria{
		Id: expected.Id,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSaveErrDuplicateKey() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	err := suite.sut.Save(random)

	suite.hashing.AssertExpectations(suite.T())

	var actual *errors.ErrAlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrAlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDuplicateKeyError",
		What:  "Id already registered",
		Why: errors.Meta{
			"Field": "Id",
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UserTestSuite) TestVerify() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Verify(random.Id))

	criteria := &repository.UserSearchCriteria{
		Id: random.Id,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.True(actual.Verified.Value)
}

func (suite *UserTestSuite) TestUpdate() {
	expected := user.Random()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NoError(suite.sut.Save(expected))

	expected.Password = user.PasswordWithValidValue()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NoError(suite.sut.Update(expected))

	suite.hashing.AssertExpectations(suite.T())

	criteria := &repository.UserSearchCriteria{
		Id: expected.Id,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestDelete() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Delete(random.Id))

	criteria := &repository.UserSearchCriteria{
		Id: random.Id,
	}

	_, err := suite.sut.Search(criteria)

	suite.Error(err)
}

func (suite *UserTestSuite) TestSearch() {
	expected := user.Random()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NoError(suite.sut.Save(expected))

	criteria := &repository.UserSearchCriteria{
		Id: expected.Id,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSearchErrDocumentNotFound() {
	random := user.Random()

	criteria := &repository.UserSearchCriteria{
		Id: random.Id,
	}

	_, err := suite.sut.Search(criteria)

	var actual *errors.ErrNotExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrNotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", random.Id.Value),
		Why: errors.Meta{
			"Index": random.Id.Value,
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
