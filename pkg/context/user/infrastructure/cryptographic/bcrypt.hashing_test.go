package cryptographic_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
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
	password, _ := valueobj.RandomPassword()

	plain := password.Value()

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	suite.NotEqual(plain, hashed)
}

func (suite *UserBcryptHashingTestSuite) TestIsNotEqual() {
	password, _ := valueobj.RandomPassword()

	plain := password.Value()

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	isNotEqual := suite.sut.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestIntegrationUserBcryptHashingSuite(t *testing.T) {
	suite.Run(t, new(UserBcryptHashingTestSuite))
}
