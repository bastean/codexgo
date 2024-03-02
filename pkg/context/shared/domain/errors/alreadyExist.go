package errors

type AlreadyExist struct {
	Message string
}

func (error AlreadyExist) Error() string {
	return error.Message
}
