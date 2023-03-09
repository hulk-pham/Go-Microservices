package common

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type Claims struct {
	UserClaim
	jwt.RegisteredClaims
}

func GenerateJWT(userClaim UserClaim) (string, error) {
	config := AppConfig()

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		UserClaim: userClaim,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(token string) (userClain UserClaim, err error) {
	claims := &Claims{}
	config := AppConfig()

	fmt.Print(token)

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})

	userClain = claims.UserClaim

	if err != nil {
		return userClain, err
	}

	if !tkn.Valid {
		return userClain, fmt.Errorf("invalid token")
	}

	return
}
