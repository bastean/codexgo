package bcrypt_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/bcrypt"
	"github.com/stretchr/testify/suite"
)

type BcryptHashingTestSuite struct {
	suite.Suite
	sut model.Hashing
}

func (suite *BcryptHashingTestSuite) SetupTest() {
	suite.sut = new(bcrypt.Bcrypt)
}

func (suite *BcryptHashingTestSuite) TestHash() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	suite.NotEqual(plain, hashed)
}

func (suite *BcryptHashingTestSuite) TestIsNotEqual() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	isNotEqual := suite.sut.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestIntegrationBcryptHashingSuite(t *testing.T) {
	suite.Run(t, new(BcryptHashingTestSuite))
}
