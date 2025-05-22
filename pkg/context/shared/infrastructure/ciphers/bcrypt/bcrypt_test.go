package bcrypt_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers/bcrypt"
)

type BcryptTestSuite struct {
	ciphers.HasherSuite
}

func (s *BcryptTestSuite) SetupSuite() {
	s.HasherSuite.SUT = new(bcrypt.Bcrypt)
}

func TestIntegrationBcryptSuite(t *testing.T) {
	suite.Run(t, new(BcryptTestSuite))
}
