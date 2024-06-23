package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type Token interface {
	JwtToken(secret string, claims jwt.MapClaims) (string, error)
}

type token struct{}

func NewToken() Token {
	return token{}
}

func (t token) JwtToken(secret string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
