package ciphers

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type HasherSuite struct {
	suite.Suite
	SUT roles.Hasher
}

func (s *HasherSuite) TestHash() {
	plain := services.Create.LoremIpsumWord()

	hashed, err := s.SUT.Hash(plain)

	s.NoError(err)

	s.NotEqual(plain, hashed)
}

func (s *HasherSuite) TestCompare() {
	hashed := services.Create.LoremIpsumWord()

	plain := services.Create.LoremIpsumWord()

	err := s.SUT.Compare(hashed, plain)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Compare",
		What:  "Do not match",
	}}

	s.Equal(expected, actual)
}
