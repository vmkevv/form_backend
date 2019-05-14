package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MakeRes function to build the json response for requests
func MakeRes(ok bool, content interface{}) map[string]interface{} {
	return gin.H{
		"success": ok,
		"content": content,
	}
}

// MakeR function most simplified
func MakeR(c *gin.Context, httpStatus int, content interface{}) {
	ok := false
	if httpStatus == http.StatusOK {
		ok = true
	}
	c.JSON(
		httpStatus,
		MakeRes(ok, content),
	)
}
