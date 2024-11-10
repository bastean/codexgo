package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
)

var (
	CreateHandler *create.Handler
	ReadHandler   *read.Handler
	UpdateHandler *update.Handler
	DeleteHandler *delete.Handler
	VerifyHandler *verify.Handler
	LoginHandler  *login.Handler
)

var (
	CreateCommandKey = create.CommandKey
	UpdateCommandKey = update.CommandKey
	DeleteCommandKey = delete.CommandKey
	VerifyCommandKey = verify.CommandKey
)

type (
	CreateCommandAttributes = create.CommandAttributes
	UpdateCommandAttributes = update.CommandAttributes
	DeleteCommandAttributes = delete.CommandAttributes
	VerifyCommandAttributes = verify.CommandAttributes
)

type (
	CreateCommandMeta = create.CommandMeta
	UpdateCommandMeta = update.CommandMeta
	DeleteCommandMeta = delete.CommandMeta
	VerifyCommandMeta = verify.CommandMeta
)

type (
	ReadQuery  = read.Query
	LoginQuery = login.Query
)

type (
	ReadResponse  = read.Response
	LoginResponse = login.Response
)
