package event

type event struct {
	PutAuthorization    string
	DeleteAuthorization string
}

var Client = event{PutAuthorization: "codexgo:put-authorization", DeleteAuthorization: "codexgo:delete-authorization"}
