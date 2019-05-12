package main

import (
	"net/http"

	"form-backend/handlers"
	"form-backend/middleware"

	"form-backend/db"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello World",
	})
}

func welcome(c *gin.Context) {
	id, _ := c.Get("id")
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{"id": id, "email": email})
}

func main() {
	db.DBCon = db.Connect()
	r := gin.Default()
	r.Use(middleware.WithCors())
	public := r.Group("/api")
	public.POST("/login", handlers.Login)

	// user := r.Group("/api")

	admin := r.Group("/api")
	admin.Use(middleware.Auth())
	admin.POST("/user", handlers.NewUser)
	admin.GET("/welcome", welcome)
	r.GET("/", hello)
	r.Run() // listen and serve on 0.0.0.0:8080
}
