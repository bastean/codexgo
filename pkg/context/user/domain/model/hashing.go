package model

type Hashing interface {
	Hash(plain string) string
	IsNotEqual(hashed, plain string) bool
}
