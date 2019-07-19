package handlers

import (
	"form-backend/db"
	"form-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetFormList get all user list
func GetFormList(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	user := db.User{}
	user.ID = userID
	if err := user.GetByID(); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	forms, err := user.GetFormsList()
	if err != nil {
		utils.MakeR(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"user":  user,
			"forms": forms,
		},
	)
}
