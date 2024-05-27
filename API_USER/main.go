package main

import (
	"api/controllers"
	"api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.Create)
	r.PUT("/posts/:id", controllers.Update)
	r.GET("/posts", controllers.READALL)
	r.GET("/posts/:id", controllers.READONE)
	r.DELETE("/posts/:id", controllers.Deletes)
	r.Run() 
}
