package persistence_test

import (
	"testing"

	database "github.com/bastean/codexgo/context/pkg/shared/infrastructure/persistence"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/context/pkg/user/infrastructure/persistence"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/cryptographic"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type MongoUserRepositoryTestSuite struct {
	suite.Suite
	userCollection persistence.UserCollection
	hashing        *cryptographic.UserHashingMock
}

func (suite *MongoUserRepositoryTestSuite) SetupTest() {
	database := database.NewMongoDatabase()
	suite.userCollection = *persistence.NewUserCollection(database, suite.hashing)
}

func (suite *MongoUserRepositoryTestSuite) TestSave() {
	user := create.Random()
	suite.NotPanics(func() { suite.userCollection.Save(user) })
}

func (suite *MongoUserRepositoryTestSuite) TestSaveDuplicate() {
	user := create.Random()

	suite.NotPanics(func() { suite.userCollection.Save(user) })

	suite.Panics(func() { suite.userCollection.Save(user) })
}

func (suite *MongoUserRepositoryTestSuite) TestUpdate() {
	user := create.Random()

	suite.NotPanics(func() { suite.userCollection.Save(user) })

	suite.NotPanics(func() { suite.userCollection.Update(user) })
}

func (suite *MongoUserRepositoryTestSuite) TestDelete() {
	user := create.Random()

	suite.NotPanics(func() { suite.userCollection.Save(user) })

	suite.NotPanics(func() { suite.userCollection.Delete(user.Id) })
}

func (suite *MongoUserRepositoryTestSuite) TestSearch() {
	user := create.Random()

	suite.NotPanics(func() { suite.userCollection.Save(user) })

	filter := repository.Filter{Email: user.Email}

	found := suite.userCollection.Search(filter)

	suite.EqualValues(user, found)
}

func TestMongoUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoUserRepositoryTestSuite))
}
