package controllers

import (
	"github.com/dortlii/simple-api/initializers"
	"github.com/dortlii/simple-api/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// repond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// repond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	// Get data from off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}
