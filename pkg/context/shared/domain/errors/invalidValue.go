package errors

type InvalidValue struct {
	Message string
}

func (error *InvalidValue) Error() string {
	return error.Message
}

func NewInvalidValue(message string) *InvalidValue {
	return &InvalidValue{
		Message: message,
	}
}
