package roles

type Hasher interface {
	Hash(plain string) (string, error)
	Compare(hashed, plain string) error
}
