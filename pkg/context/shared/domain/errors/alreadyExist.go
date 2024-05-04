package errors

type AlreadyExist struct {
	Message string
}

func (error *AlreadyExist) Error() string {
	return error.Message
}

func NewAlreadyExist(message string) *AlreadyExist {
	return &AlreadyExist{
		Message: message,
	}
}
