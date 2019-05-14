package utils

import (
	"form-backend/structs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenToken generates a token, based in id and email
func GenToken(id int, email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &structs.Claims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(structs.JwtKey)
	return tokenString, err
}
