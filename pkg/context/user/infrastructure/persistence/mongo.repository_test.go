package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type MongoRepositoryTestSuite struct {
	suite.Suite
	sut     model.Repository
	hashing *cryptographic.HashingMock
}

func (suite *MongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")

	databaseName := "codexgo-test"

	database, _ := persistences.NewMongoDatabase(uri, databaseName)

	collectionName := "users-test"

	suite.hashing = new(cryptographic.HashingMock)

	suite.sut, _ = persistence.NewMongoCollection(database, collectionName, suite.hashing)
}

func (suite *MongoRepositoryTestSuite) TestSave() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *MongoRepositoryTestSuite) TestSaveDuplicate() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	err := suite.sut.Save(user)

	suite.hashing.AssertExpectations(suite.T())

	var actual *errors.AlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleMongoDuplicateKeyError",
		What:  "already registered",
		Why: errors.Meta{
			"Field": "Id",
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *MongoRepositoryTestSuite) TestVerify() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Verify(user.Id))
}

func (suite *MongoRepositoryTestSuite) TestUpdate() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	password := valueobj.PasswordWithValidValue()

	user.Password = password

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Update(user))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *MongoRepositoryTestSuite) TestDelete() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Id))
}

func (suite *MongoRepositoryTestSuite) TestSearch() {
	expected := aggregate.RandomUser()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value()).Return(expected.Password.Value())

	suite.NoError(suite.sut.Save(expected))

	criteria := &model.RepositorySearchCriteria{
		Id: expected.Id,
	}

	user, err := suite.sut.Search(criteria)

	suite.NoError(err)

	actual := user

	suite.Equal(expected, actual)
}

func TestIntegrationMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoRepositoryTestSuite))
}
