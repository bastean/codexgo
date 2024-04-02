package verify

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id *valueObject.Id) {
	userRegistered := verify.Repository.Search(model.RepositorySearchFilter{Id: id})

	// TODO?: if userRegistered.Verified.Value { return }

	userRegistered.Verified = valueObject.NewVerified(true)

	// TODO?: userRegistered.Password = nil
	userRegistered.Password = nil

	verify.Repository.Update(userRegistered)
}

func NewVerify(repository model.Repository) *Verify {
	return &Verify{
		Repository: repository,
	}
}
