package handlers

import (
	"fmt"
	"form-backend/db"
	"form-backend/structs"
	"form-backend/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetFormDoc get a doc form
func GetFormDoc(c *gin.Context) {
	nro := c.Param("nro")
	form := db.FormDoc{}
	user := db.User{}
	if err := form.GetByNro(nro); err != nil {
		fmt.Println(err.Error())
		utils.MakeR(c, http.StatusBadRequest, "No existe el form nro "+nro)
		return
	}
	user.ID = form.UserID
	if err := user.GetByID(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"form": form,
			"user": user,
		},
	)
}

// NewFormDoc saves new doc form
func NewFormDoc(c *gin.Context) {
	tokenString, _ := c.Get("tokenString")
	claims := &structs.Claims{}
	_, err := jwt.ParseWithClaims(tokenString.(string), claims, func(tokenString *jwt.Token) (interface{}, error) {
		return structs.JwtKey, nil
	})
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	user := db.User{}
	if err := user.GetByEmail(claims.Email); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	var formDoc db.FormDoc
	if err := c.BindJSON(&formDoc); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	formDoc.UpdatedAt = time.Now()
	formDoc.CreatedAt = time.Now()
	formDoc.UserID = user.ID
	if err := formDoc.Save(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"token": c.GetString("tokenString"),
		},
	)
}
