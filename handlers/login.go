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
	admin := db.Admin{
		Email: cred.Correo,
	}
	user := db.User{
		Email: cred.Correo,
	}
	errAdmin := admin.GetByEmail(cred.Correo)
	errUser := user.GetByEmail(cred.Correo)

	if errAdmin != nil && errUser != nil {
		utils.MakeR(c, http.StatusBadRequest, "Correo electrónico no válido")
		return
	}
	if errUser == nil {
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
				"type":  "user",
				"user":  user,
			},
		)
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(cred.Password)); err != nil {
			utils.MakeR(c, http.StatusBadRequest, "Password incorrecto, vuelva a intentarlo.")
			return
		}
		tokenString, err := utils.GenToken(admin.ID, admin.Email)
		if err != nil {
			utils.MakeR(c, http.StatusInternalServerError, "Error al generar el token")
			return
		}
		utils.MakeR(
			c,
			http.StatusOK,
			gin.H{
				"token": tokenString,
				"type":  "admin",
				"user":  admin,
			},
		)
	}
}
