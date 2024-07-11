package collection_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/collection"
	"github.com/stretchr/testify/suite"
)

type MongoRepositoryTestSuite struct {
	suite.Suite
	sut     model.Repository
	hashing *cryptographic.HashingMock
}

func (suite *MongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_MONGODB_URI")

	name := os.Getenv("DATABASE_MONGODB_NAME")

	database, _ := mongodb.New(uri, name)

	name = "users-test"

	suite.hashing = new(cryptographic.HashingMock)

	suite.sut, _ = collection.NewUser(database, name, suite.hashing)
}

func (suite *MongoRepositoryTestSuite) TestSave() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *MongoRepositoryTestSuite) TestSaveDuplicate() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	err := suite.sut.Save(random)

	suite.hashing.AssertExpectations(suite.T())

	var actual *errors.AlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleDuplicateKeyError",
		What:  "already registered",
		Why: errors.Meta{
			"Field": "Id",
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *MongoRepositoryTestSuite) TestVerify() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Verify(random.Id))
}

func (suite *MongoRepositoryTestSuite) TestUpdate() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	random.Password = user.PasswordWithValidValue()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Update(random))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *MongoRepositoryTestSuite) TestDelete() {
	random := user.Random()

	suite.hashing.On("Hash", random.Password.Value).Return(random.Password.Value)

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Delete(random.Id))
}

func (suite *MongoRepositoryTestSuite) TestSearch() {
	expected := user.Random()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NoError(suite.sut.Save(expected))

	criteria := &model.RepositorySearchCriteria{
		Id: expected.Id,
	}

	actual, err := suite.sut.Search(criteria)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func TestIntegrationMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoRepositoryTestSuite))
}
