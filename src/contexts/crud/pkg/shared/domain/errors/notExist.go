package errors

type NotExist struct {
	Message string
}

func (error NotExist) Error() string {
	return error.Message
}
