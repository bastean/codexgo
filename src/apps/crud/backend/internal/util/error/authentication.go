package error

func AuthenticationMissing(field string) []Error {
	return []Error{{Field: field, Message: "Missing"}}
}
