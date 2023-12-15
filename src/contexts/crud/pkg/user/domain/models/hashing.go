package models

type Hashing interface {
	Hash(plain string) string
	IsNotEqual(plain, hashed string) bool
}
