package cryptographic_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/stretchr/testify/suite"
)

type UserBcryptHashingTestSuite struct {
	suite.Suite
	sut model.Hashing
}

func (suite *UserBcryptHashingTestSuite) SetupTest() {
	suite.sut = cryptographic.NewUserBcryptHashing()
}

func (suite *UserBcryptHashingTestSuite) TestHash() {
	plain := valueObjectMother.RandomPassword().Value

	hashed := suite.sut.Hash(plain)

	suite.NotEqual(plain, hashed)
}

func (suite *UserBcryptHashingTestSuite) TestIsNotEqual() {
	plain := valueObjectMother.RandomPassword().Value

	hashed := suite.sut.Hash(plain)

	isNotEqual := suite.sut.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestUserBcryptHashingSuite(t *testing.T) {
	suite.Run(t, new(UserBcryptHashingTestSuite))
}
