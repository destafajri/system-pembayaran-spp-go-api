package helper

import (
	"log"
	"time"

	middleware"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(name, phone, role string) (string, error) {
	//proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 300)
	claims := middleware.JWTClaim{
		Name:   name,
		Phone: phone,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "github.com/destafajri",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//medeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//signed token
	token, err := tokenAlgo.SignedString(middleware.JWT_SECRET_KEY)
	if err != nil {
		log.Println("signed token error")
		return "", err
	}

	return token, nil
}