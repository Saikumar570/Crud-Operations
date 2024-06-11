package controllers

import (
	"api/models"
	"api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Repo *repositories.PostRepository
}

func NewPostController(repo *repositories.PostRepository) *PostController {
	return &PostController{Repo: repo}
}

func (pc *PostController) Create(c *gin.Context) {
	var post models.Post
	if err := c.Bind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := pc.Repo.Create(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) ReadAll(c *gin.Context) {
	var posts []models.Post
	if err := pc.Repo.FindAll(&posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (pc *PostController) ReadOne(c *gin.Context) {
	id := c.Param("post_id")
	var post models.Post
	if err := pc.Repo.FindByID(id, &post); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) Update(c *gin.Context) {
	id := c.Param("post_id")
	var post models.Post
	if err := pc.Repo.FindByID(id, &post); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := c.Bind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := pc.Repo.Update(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) Delete(c *gin.Context) {
	id := c.Param("post_id")
	if err := pc.Repo.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete post"})
		return
	}
	c.Status(http.StatusOK)
}

type CommentController struct {
	Repo *repositories.CommentRepository
}

func NewCommentController(repo *repositories.CommentRepository) *CommentController {
	return &CommentController{Repo: repo}
}

func (cc *CommentController) Create(c *gin.Context) {
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	comment.PostID = uint(postID)
	if err := cc.Repo.Create(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func (cc *CommentController) ReadAllByPostID(c *gin.Context) {
	postID := c.Param("post_id")
	var comments []models.Comment
	if err := cc.Repo.FindAllByPostID(postID, &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (cc *CommentController) ReadOne(c *gin.Context) {
	id := c.Param("comment_id")
	var comment models.Comment
	if err := cc.Repo.FindByID(id, &comment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func (cc *CommentController) Update(c *gin.Context) {
	id := c.Param("comment_id")
	var comment models.Comment
	if err := cc.Repo.FindByID(id, &comment); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := c.Bind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := cc.Repo.Update(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func (cc *CommentController) Delete(c *gin.Context) {
	id := c.Param("comment_id")
	if err := cc.Repo.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete comment"})
		return
	}
	c.Status(http.StatusOK)
}
