package middlewares

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// helper variable
var (
	JWT_SECRET_KEY = []byte(os.Getenv("KEY_JWT"))
)

// claims struct
type JWTClaim struct {
	Name  string
	Phone string
	Role  string
	jwt.RegisteredClaims
}

func GetRole(claims *JWTClaim) string{
	return claims.Role
}

func GetClaims(claims *JWTClaim) *JWTClaim{
	return claims
}