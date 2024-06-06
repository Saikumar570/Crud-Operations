package controllers

import (
	"api/models"
	"api/repositories"

	"github.com/gin-gonic/gin"
)

var postRepository *repositories.Repository

func PostRepository(repo *repositories.Repository) {
	postRepository = repo
}

func Create(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := postRepository.Create(&post); err != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func ReadAll(c *gin.Context) {
	var posts []models.Post
	if err := postRepository.FindAll(&posts); err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve posts"})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

func ReadOne(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := postRepository.FindByID(id, &post); err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := postRepository.FindByID(id, &post); err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	if err := c.Bind(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := postRepository.Update(&post); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func Deletes(c *gin.Context) {
	id := c.Param("id")

	if err := postRepository.Delete(id); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete post"})
		return
	}

	c.Status(200)
}
