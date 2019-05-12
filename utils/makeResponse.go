package utils

import "github.com/gin-gonic/gin"

// MakeRes function to build the json response for requests
func MakeRes(ok bool, content interface{}) map[string]interface{} {
	return gin.H{
		"success": ok,
		"content": content,
	}
}
