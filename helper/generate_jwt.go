package helper

import (
	"log"
	"time"

	middleware "github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(id, username, role string) (string, error) {
	// create jwt token
	expTime := time.Now().Add(time.Minute * 300)
	claims := middleware.JWTClaim{
		ID:       id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/destafajri",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// declare algoritm for signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(middleware.JWT_SECRET_KEY)
	if err != nil {
		log.Println("signed token error")
		return "", err
	}

	return token, nil
}
