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

func (s *BcryptTestSuite) SetupTest() {
	s.sut = new(bcrypt.Bcrypt)
}

func (s *BcryptTestSuite) TestHash() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := s.sut.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *BcryptTestSuite) TestIsNotEqual() {
	password := user.PasswordWithValidValue()

	plain := password.Value

	hashed, err := s.sut.Hash(plain)

	s.NoError(err)

	isNotEqual := s.sut.IsNotEqual(hashed, plain)

	s.False(isNotEqual)
}

func TestIntegrationBcryptSuite(t *testing.T) {
	suite.Run(t, new(BcryptTestSuite))
}
