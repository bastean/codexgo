package cryptographic_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	create "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
	"github.com/stretchr/testify/suite"
)

type BcryptUserHashingTestSuite struct {
	suite.Suite
	bcrypt cryptographic.Bcrypt
}

func (suite *BcryptUserHashingTestSuite) SetupTest() {}

func (suite *BcryptUserHashingTestSuite) TestHash() {
	plain := create.RandomPassword().Value

	hashed := suite.bcrypt.Hash(plain)

	suite.NotEqual(plain, hashed)
}

func (suite *BcryptUserHashingTestSuite) TestIsNotEqual() {
	plain := create.RandomPassword().Value

	hashed := suite.bcrypt.Hash(plain)

	isNotEqual := suite.bcrypt.IsNotEqual(hashed, plain)

	suite.False(isNotEqual)
}

func TestBcryptUserHashingSuite(t *testing.T) {
	suite.Run(t, new(BcryptUserHashingTestSuite))
}
