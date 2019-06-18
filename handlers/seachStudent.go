package handlers

import (
	"form-backend/db"
	"form-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentCredentials struct {
	Ci string `json:"ci"`
}

// SearchStudent searchs a student by ci
func SearchStudent(c *gin.Context) {
	var cred studentCredentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		utils.MakeR(c, http.StatusInternalServerError, "Error al parsear JSON")
		return
	}
	student := db.Student{
		Ci: cred.Ci,
	}
	if err := student.GetByCi(cred.Ci); err != nil {
		utils.MakeR(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.MakeR(
		c,
		http.StatusOK,
		gin.H{
			"student": student,
		},
	)
}
