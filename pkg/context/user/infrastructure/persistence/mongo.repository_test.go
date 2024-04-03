package persistence_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence/database"
	aggregateMother "github.com/bastean/codexgo/pkg/context/user/domain/aggregate/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserMongoRepositoryTestSuite struct {
	suite.Suite
	sut     model.Repository
	hashing *cryptographicMock.HashingMock
}

func (suite *UserMongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")
	databaseName := "codexgo-test"
	database := database.NewMongoDatabase(uri, databaseName)
	collectionName := "users-test"
	suite.hashing = cryptographicMock.NewHashingMock()
	suite.sut = persistence.NewUserMongoRepository(database, collectionName, suite.hashing)
}

func (suite *UserMongoRepositoryTestSuite) TestSave() {
	user := aggregateMother.Random()

	suite.hashing.On("Hash", user.Password.Value).Return(user.Password.Value)

	suite.NotPanics(func() { suite.sut.Save(user) })

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *UserMongoRepositoryTestSuite) TestSaveDuplicate() {
	user := aggregateMother.Random()

	suite.hashing.On("Hash", user.Password.Value).Return(user.Password.Value)

	suite.NotPanics(func() { suite.sut.Save(user) })

	suite.Panics(func() { suite.sut.Save(user) })
}

func (suite *UserMongoRepositoryTestSuite) TestUpdate() {
	user := aggregateMother.Random()

	suite.hashing.On("Hash", user.Password.Value).Return(user.Password.Value)

	suite.NotPanics(func() { suite.sut.Save(user) })

	user.Password = valueObjectMother.RandomPassword()

	suite.hashing.On("Hash", user.Password.Value).Return(user.Password.Value)

	suite.NotPanics(func() { suite.sut.Update(user) })

	suite.hashing.AssertExpectations(suite.T())
}

func (suite *UserMongoRepositoryTestSuite) TestDelete() {
	user := aggregateMother.Random()

	suite.hashing.On("Hash", user.Password.Value).Return(user.Password.Value)

	suite.NotPanics(func() { suite.sut.Save(user) })

	suite.NotPanics(func() { suite.sut.Delete(user.Id) })
}

func (suite *UserMongoRepositoryTestSuite) TestSearch() {
	expected := aggregateMother.Random()

	expected.PullMessages()

	suite.hashing.On("Hash", expected.Password.Value).Return(expected.Password.Value)

	suite.NotPanics(func() { suite.sut.Save(expected) })

	filter := model.RepositorySearchFilter{Email: expected.Email}

	actual := suite.sut.Search(filter)

	suite.EqualValues(expected, actual)
}

func TestIntegrationUserMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserMongoRepositoryTestSuite))
}
