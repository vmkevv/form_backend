package middleware

import (
	"fmt"
	"form-backend/structs"
	"form-backend/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth: handles token validation
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, "No se reconoce el token en los headers."))
			return
		}
		claims := &structs.Claims{}
		tokenStruct, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return structs.JwtKey, nil
		})
		if err != nil {
			if strings.Contains(err.Error(), "token is expired by") {
				c.Abort()
				c.JSON(http.StatusBadRequest, utils.MakeRes(false, "Ha estado inactivo por más de 25 minutos, por favor ingrese de nuevo."))
				return
			}
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, err.Error()))
			return
		}
		if !tokenStruct.Valid {
			c.Abort()
			c.JSON(http.StatusInternalServerError, utils.MakeRes(false, "Token no valido."))
			return
		}
		fmt.Println(claims.ID)
		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
	}
}
