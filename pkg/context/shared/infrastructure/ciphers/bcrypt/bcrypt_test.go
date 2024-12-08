package bcrypt_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers/bcrypt"
)

type BcryptTestSuite struct {
	ciphers.HasherSuite
}

func (s *BcryptTestSuite) SetupTest() {
	s.HasherSuite.SUT = new(bcrypt.Bcrypt)
}

func TestIntegrationBcryptSuite(t *testing.T) {
	suite.Run(t, new(BcryptTestSuite))
}
