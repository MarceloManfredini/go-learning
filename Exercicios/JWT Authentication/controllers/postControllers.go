package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Parse request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Create a post
	/* 	post := models.Post{
		Title: "My last post",
		Body:  "This is the body of my last post",
	} */
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// Return them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id from url
	id := c.Param("id")

	// Get post
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get id from url
	id := c.Param("id")
	// Get post
	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Post not found",
		})
		return
	}
	// Update it
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post.Title = body.Title
	post.Body = body.Body
	initializers.DB.Save(&post)
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
