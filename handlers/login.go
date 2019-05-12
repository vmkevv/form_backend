package handlers

import (
	"form-backend/db"
	"form-backend/structs"
	"form-backend/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

// Login func to login
func Login(c *gin.Context) {
	var cred credentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.MakeRes(false, "Error al parsear JSON"),
		)
		return
	}
	admin := db.Admin{
		Email: cred.Correo,
	}
	err := admin.GetByEmail(cred.Correo)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.MakeRes(false, "Correo electrónico no válido"),
		)
		return
	}
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.MakeRes(false, "No se pudo encriptar el password"),
		)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(cred.Password)); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.MakeRes(false, "Password incorrecto, vuelva a intentarlo."),
		)
		return
	}
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &structs.Claims{
		ID:    admin.ID,
		Email: admin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(structs.JwtKey)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.MakeRes(false, "Error al generar el token."),
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"token": tokenString,
			"admin": admin,
		},
	)
}
