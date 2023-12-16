package persistence

import (
	"fmt"

	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/davecgh/go-spew/spew"
)

type Mongo struct{}

func (mongo Mongo) Save(user *aggregate.User) {
	fmt.Println("SAVE")
	spew.Dump(user)
}

func (mongo Mongo) Update(user *aggregate.User) {
	fmt.Println("UPDATE")
	spew.Dump(user)
}

func (mongo Mongo) Delete(email *sharedVO.Email) {
	fmt.Println("DELETE")
	spew.Dump(email)
}

func (mongo Mongo) Search(email *sharedVO.Email) (*aggregate.User, error) {
	fmt.Println("SEARCH")
	spew.Dump(email)
	return &aggregate.User{}, nil
}
