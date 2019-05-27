package handlers

import (
	"form-backend/db"
	"form-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NewFormEst saves new form
func NewFormEst(c *gin.Context) {
	var formEst db.FormEst
	err := c.BindJSON(&formEst)
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "No se pudo parsear")
		return
	}
	formEst.UpdatedAt = time.Now()
	formEst.CreatedAt = time.Now()
	if err := formEst.Save(db.DBCon); err != nil {
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
