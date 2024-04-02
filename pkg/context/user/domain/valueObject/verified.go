package valueObject

type Verified struct {
	Value bool
}

func NewVerified(verified bool) *Verified {
	verifiedVO := &Verified{Value: verified}

	return verifiedVO
}
