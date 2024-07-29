package bcrypt_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic/bcrypt"
	"github.com/stretchr/testify/suite"
)

type BcryptTestSuite struct {
	suite.Suite
	sut hashing.Hashing
}

func (suite *BcryptTestSuite) SetupTest() {
	suite.sut = new(bcrypt.Bcrypt)
}

func (suite *BcryptTestSuite) TestHash() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	suite.NotEqual(plain, hashed)
}

func (suite *BcryptTestSuite) TestIsNotEqual() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := suite.sut.Hash(plain)

	suite.NoError(err)

	isNotEqual := suite.sut.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestIntegrationBcryptSuite(t *testing.T) {
	suite.Run(t, new(BcryptTestSuite))
}
