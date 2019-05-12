package structs

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtKey key for JWT
var JwtKey = []byte("serverkey")

// Claims for make the token
type Claims struct {
	Email string `json:"correo"`
	ID    int    `json:"id"`
	jwt.StandardClaims
}
