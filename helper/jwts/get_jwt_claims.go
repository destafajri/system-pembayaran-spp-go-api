package jwts

import (
	"log"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/middlewares"
	"github.com/golang-jwt/jwt/v4"
)


func GetID(token string) (string, error) {
	claims := &middlewares.JWTClaim{}

	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(middlewares.JWT_SECRET_KEY), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return claims.ID, nil
}

func GetUsername(token string) (string, error) {
	claims := &middlewares.JWTClaim{}

	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(middlewares.JWT_SECRET_KEY), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return claims.Username, nil
}

func GetRole(token string) (string, error) {
	claims := &middlewares.JWTClaim{}

	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(middlewares.JWT_SECRET_KEY), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	return claims.Role, nil
}

func GetClaims(token string) (*middlewares.JWTClaim, error) {
	claims := &middlewares.JWTClaim{}
	
	// parsing token jwt
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(middlewares.JWT_SECRET_KEY), nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return claims, nil
}
