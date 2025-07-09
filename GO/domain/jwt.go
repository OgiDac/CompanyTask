package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
