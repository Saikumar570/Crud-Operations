package controllers

import (
	"api/initializers"
	"api/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	//get data i.e. to see
	var body struct {
		Name       string
		Employeeid int
		Email      string
	}

	c.Bind(&body)

	// create post variable to store
	post := models.Post{Name: body.Name, Employeeid: body.Employeeid, Email: body.Email}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func READALL(c *gin.Context) {
	//get the all posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func READONE(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	//get the post of particular id
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Update(c *gin.Context) {
	//get id
	id := c.Param("id")
	//get data
	var body struct {
		Name       string
		Employeeid int
		Email      string
	}
	c.Bind(&body)
	// find it
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Name:       body.Name,
		Employeeid: body.Employeeid,
		Email:      body.Email,
	})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func Deletes(c *gin.Context) {
	//get id
	id := c.Param("id")
	//delete posts
	initializers.DB.Delete(&models.Post{}, id)
	//respond
	c.Status(200)
}
