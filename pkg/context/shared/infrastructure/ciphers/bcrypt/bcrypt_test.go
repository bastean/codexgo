package bcrypt_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers/bcrypt"
)

type BcryptTestSuite struct {
	suite.Suite
	sut hashes.Hashing
}

func (s *BcryptTestSuite) SetupTest() {
	s.sut = new(bcrypt.Bcrypt)
}

func (s *BcryptTestSuite) TestHash() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.sut.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *BcryptTestSuite) TestIsNotEqual() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.sut.Hash(plain)

	s.NoError(err)

	s.False(s.sut.IsNotEqual(hashed, plain))
}

func TestIntegrationBcryptSuite(t *testing.T) {
	suite.Run(t, new(BcryptTestSuite))
}
