package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"log"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var bodynya struct {
		Title2 string
		Body2  string
	}

	c.Bind(&bodynya)

	post := models.Post{Title: bodynya.Title2, Body: bodynya.Body2}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostIndexId(c *gin.Context) {

	id := c.Param("id")

	var posts models.Post
	initializers.DB.First(&posts, id)

	c.JSON(200, gin.H{
		"post": posts,
	})

}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var pos models.Post
	initializers.DB.First(&pos, id)

	result := initializers.DB.Model(&pos).Updates(models.Post{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		log.Fatal("Gagal")
		c.JSON(500, gin.H{
			"post": "Gagal",
		})

	}

	c.JSON(200, gin.H{
		"post": pos,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
