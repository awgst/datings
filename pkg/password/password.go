package password

import (
	"crypto/sha512"
	"encoding/base64"
)

type Password interface {
	Hash(password string) string
	Compare(hashedPassword string, password string) bool
}

type password struct {
}

func (p password) Hash(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))

	hashedBytes := hasher.Sum(nil)
	hashedBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashedBase64
}

func (p password) Compare(hashedPassword string, password string) bool {
	return p.Hash(password) == hashedPassword
}

func NewPassword() Password {
	return &password{}
}
