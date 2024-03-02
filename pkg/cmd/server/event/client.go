package event

type event struct {
	PutAuthorization    string
	DeleteAuthorization string
}

var Client = event{"codexgo:put-authorization", "codexgo:delete-authorization"}
