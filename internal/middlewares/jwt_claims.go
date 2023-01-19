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
	ID       string
	Username string
	Role     string
	jwt.RegisteredClaims
}