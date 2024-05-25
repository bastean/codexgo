package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/spersistence"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserMongoRepositoryTestSuite struct {
	suite.Suite
	sut     model.Repository
	hashing *cryptographic.HashingMock
}

func (suite *UserMongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")

	databaseName := "codexgo-test"

	database, _ := spersistence.NewMongoDatabase(uri, databaseName)

	collectionName := "users-test"

	suite.hashing = new(cryptographic.HashingMock)

	suite.sut, _ = persistence.NewMongoCollection(database, collectionName, suite.hashing)
}

func (suite *UserMongoRepositoryTestSuite) TestSave() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *UserMongoRepositoryTestSuite) TestSaveDuplicate() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.Error(suite.sut.Save(user))
}

func (suite *UserMongoRepositoryTestSuite) TestUpdate() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	password, _ := valueobj.RandomPassword()

	user.Password = password

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Update(user))

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *UserMongoRepositoryTestSuite) TestDelete() {
	user := aggregate.RandomUser()

	suite.hashing.On("Hash", user.Password.Value()).Return(user.Password.Value())

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Id))
}

func (suite *UserMongoRepositoryTestSuite) TestSearch() {
	expected := aggregate.RandomUser()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value()).Return(expected.Password.Value())

	suite.NoError(suite.sut.Save(expected))

	filter := model.RepositorySearchCriteria{
		Email: expected.Email,
	}

	user, err := suite.sut.Search(filter)

	suite.NoError(err)

	actual := user

	suite.EqualValues(expected, actual)
}

func TestIntegrationUserMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserMongoRepositoryTestSuite))
}
