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
	public.POST("/student", handlers.SearchStudent)
	public.GET("/user/forms/:id", handlers.GetFormList)

	admin := r.Group("/api")
	admin.Use(middleware.Auth())
	admin.POST("/user", handlers.NewUser)
	admin.GET("/user", handlers.UserData)
	admin.PUT("/user", handlers.UpdatePassword)
	admin.GET("/users", handlers.UserList)
	admin.POST("/user/active", handlers.Active)
	admin.POST("/user/reset", handlers.Reset)
	admin.POST("/form/form-est", handlers.NewFormEst)
	admin.GET("/form/form-est/:nro", handlers.GetFormEst)
	admin.POST("/form/form-pro", handlers.NewFormPro)
	admin.GET("/form/form-pro/:nro", handlers.GetFormPro)
	admin.POST("/form/form-pre", handlers.NewFormPre)
	admin.GET("/form/form-pre/:nro", handlers.GetFormPre)
	admin.POST("/form/form-doc", handlers.NewFormDoc)
	admin.GET("/form/form-doc/:nro", handlers.GetFormDoc)

	r.GET("/", hello)
	r.Run(":4000") // listen and serve on 0.0.0.0:8080
}
