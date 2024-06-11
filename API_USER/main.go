package main

import (
	"api/controllers"
	"api/initializers"
	"api/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := initializers.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repo := repositories.NewRepository(initializers.DB)
	postController := controllers.NewPostController(repo.Post)
	commentController := controllers.NewCommentController(repo.Comment)
	r := gin.Default()

	r.POST("/posts", postController.Create)
	r.GET("/posts", postController.ReadAll)
	r.GET("/posts/:post_id", postController.ReadOne)
	r.PUT("/posts/:post_id", postController.Update)
	r.DELETE("/posts/:post_id", postController.Delete)

	postGroup := r.Group("/posts/:post_id")
	{
		postGroup.POST("/comments", commentController.Create)
		postGroup.GET("/comments", commentController.ReadAllByPostID)
	}
	r.GET("/comments/:comment_id", commentController.ReadOne)
	r.PUT("/comments/:comment_id", commentController.Update)
	r.DELETE("/comments/:comment_id", commentController.Delete)

	r.Run(":3000")
}
