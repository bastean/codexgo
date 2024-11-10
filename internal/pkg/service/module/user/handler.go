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

var (
	ReadQueryKey  = read.QueryKey
	LoginQueryKey = login.QueryKey
)

var (
	ReadResponseKey  = read.ResponseKey
	LoginResponseKey = login.ResponseKey
)

type (
	CreateCommandAttributes = create.CommandAttributes
	UpdateCommandAttributes = update.CommandAttributes
	DeleteCommandAttributes = delete.CommandAttributes
	VerifyCommandAttributes = verify.CommandAttributes
)

type (
	ReadQueryAttributes  = read.QueryAttributes
	LoginQueryAttributes = login.QueryAttributes
)

type (
	ReadResponseAttributes  = read.ResponseAttributes
	LoginResponseAttributes = login.ResponseAttributes
)

type (
	CreateCommandMeta = create.CommandMeta
	UpdateCommandMeta = update.CommandMeta
	DeleteCommandMeta = delete.CommandMeta
	VerifyCommandMeta = verify.CommandMeta
)

type (
	ReadQueryMeta  = read.QueryMeta
	LoginQueryMeta = login.QueryMeta
)

type (
	ReadResponseMeta  = read.ResponseMeta
	LoginResponseMeta = login.ResponseMeta
)
