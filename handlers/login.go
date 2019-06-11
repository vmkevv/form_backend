package handlers

import (
	"form-backend/db"
	"form-backend/utils"
	"net/http"

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
		utils.MakeR(c, http.StatusInternalServerError, "Error al parsear JSON")
		return
	}
	user := db.User{
		Email: cred.Correo,
	}
	errUser := user.GetByEmail(cred.Correo)

	if errUser != nil {
		utils.MakeR(c, http.StatusBadRequest, "Correo electrónico no válido")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		utils.MakeR(c, http.StatusBadRequest, "Password incorrecto, vuelva a intentarlo.")
		return
	}
	tokenString, err := utils.GenToken(user.ID, user.Email)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error al generar el token")
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"token": tokenString,
			"user":  user,
		},
	)
}
