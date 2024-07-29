package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
)

type (
	CreateCommand = create.Command
	UpdateCommand = update.Command
	DeleteCommand = delete.Command
	VerifyCommand = verify.Command
)

type (
	ReadQuery  = read.Query
	LoginQuery = login.Query
)

type (
	ReadResponse = read.Response
)

var (
	Create *create.Handler
	Read   *read.Handler
	Update *update.Handler
	Delete *delete.Handler
	Verify *verify.Handler
	Login  *login.Handler
)
