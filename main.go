package main

import (
	"form-backend/db"
	"form-backend/handlers"
	"form-backend/middleware"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello World",
	})
}

func main() {
	db.DBCon = db.Connect()
	r := gin.Default()
	r.Use(middleware.WithCors())

	public := r.Group("/api")
	public.POST("/login", handlers.Login)

	admin := r.Group("/api")
	admin.Use(middleware.Auth())
	admin.POST("/user", handlers.NewUser)
	admin.POST("/user/active", handlers.Active)
	admin.POST("/user/reset", handlers.Reset)

	r.GET("/", hello)
	r.Run() // listen and serve on 0.0.0.0:8080
}
