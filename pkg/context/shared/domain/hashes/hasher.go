package hashes

type Hasher interface {
	Hash(plain string) (string, error)
	IsNotEqual(hashed, plain string) bool
}
