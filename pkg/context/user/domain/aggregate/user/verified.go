package user

type Verified struct {
	Value bool
}

func NewVerified(value bool) (*Verified, error) {
	return &Verified{
		Value: value,
	}, nil
}
