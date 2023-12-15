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

func (mongo Mongo) Delete(id *sharedVO.Id) {
	fmt.Println("DELETE")
	spew.Dump(id)
}

func (mongo Mongo) Search(id *sharedVO.Id) (*aggregate.User, error) {
	fmt.Println("SEARCH")
	spew.Dump(id)
	return &aggregate.User{}, nil
}
