package handlers

import (
	"form-backend/db"
	"form-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Active changes active status of user
func Active(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error parseando json")
		return
	}
	err = user.ChangeActive(db.DBCon)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Ocurrio un error al actualizar estado.")
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
