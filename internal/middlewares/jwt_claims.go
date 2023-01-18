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

func GetID(token string) (string, error) {
	claims := &JWTClaim{}
	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return "", err
	}

	return claims.ID, nil
}

func GetUsername(token string) (string, error) {
	claims := &JWTClaim{}
	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return "", err
	}

	return claims.Username, nil
}

func GetRole(token string) (string, error) {
	claims := &JWTClaim{}
	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return "", err
	}

	return claims.Role, nil
}

func GetClaims(token string) (*JWTClaim, error) {
	claims := &JWTClaim{}
	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}
	return claims, nil
}
