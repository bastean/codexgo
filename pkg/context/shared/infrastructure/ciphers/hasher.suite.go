package ciphers

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type HasherSuite struct {
	suite.Suite
	SUT hashes.Hasher
}

func (s *HasherSuite) TestHash() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *HasherSuite) TestIsNotEqual() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	s.False(s.SUT.IsNotEqual(hashed, plain))
}
