package errors

type Failed struct {
	Message string
}

func (error *Failed) Error() string {
	return error.Message
}

func NewFailed(message string) *Failed {
	return &Failed{
		Message: message,
	}
}
