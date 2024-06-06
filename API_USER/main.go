package main

import (
	"api/controllers"
	"api/initializers"
	"api/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	postRepository := repositories.FunctionOfRepository(initializers.DB)
	controllers.PostRepository(postRepository)
	r := gin.Default()

	r.POST("/posts", controllers.Create)
	r.GET("/posts", controllers.ReadAll)
	r.GET("/posts/:id", controllers.ReadOne)
	r.PUT("/posts/:id", controllers.Update)
	r.DELETE("/posts/:id", controllers.Deletes)

	r.Run(":3000")
}
