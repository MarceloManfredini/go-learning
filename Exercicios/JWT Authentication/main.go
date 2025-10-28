package main

// https://www.youtube.com/watch?v=ma7rUS_vW9M&list=PL-LRDpVN2fZAluCzYNZdSCfJVQXe5ly90&index=10

import (
	"jwt-authentication/controllers"
	"jwt-authentication/initializers"
	"jwt-authentication/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
