package errors

type Failed struct {
	Message string
}

func (error Failed) Error() string {
	return error.Message
}
