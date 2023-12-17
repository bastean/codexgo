package errors

type InvalidValue struct {
	Message string
}

func (error InvalidValue) Error() string {
	return error.Message
}
