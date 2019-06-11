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
	admin.GET("/user", handlers.UserData)
	admin.PUT("/user", handlers.UpdatePassword)
	admin.GET("/users", handlers.UserList)
	admin.POST("/user/active", handlers.Active)
	admin.POST("/user/reset", handlers.Reset)
	admin.POST("/form/form-est", handlers.NewFormEst)

	r.GET("/", hello)
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}
