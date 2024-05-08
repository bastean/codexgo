package model

type Hashing interface {
	Hash(plain string) (string, error)
	IsNotEqual(hashed, plain string) bool
}
