package cryptographic

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func (hashing Bcrypt) Hash(plain string) string {
	salt := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), salt)

	if err != nil {
		panic(err.Error())
	}

	return string(bytes)
}

func (hashing Bcrypt) IsNotEqual(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err != nil
}
