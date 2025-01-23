package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateData(c *gin.Context) {

	// Inital object
	var bodynya struct {
		Title2 string
		Body2  string
	}

	c.Bind(&bodynya)

	// Create object from struct
	post := models.Post{Title: bodynya.Title2, Body: bodynya.Body2}

	// Save data
	result := initializers.DB.Create(&post)

	// Response when geting error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Response data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetListData(c *gin.Context) {

	// Initial object list
	var posts []models.Post

	// Find to database
	initializers.DB.Find(&posts)

	// Response data
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func GetDataById(c *gin.Context) {

	// Set Parameter
	id := c.Param("id")

	// Inital object
	var posts models.Post

	// Find to database with id
	initializers.DB.First(&posts, id)

	// Response data
	c.JSON(200, gin.H{
		"post": posts,
	})

}

func UpdateDataById(c *gin.Context) {

	// Set Parameter
	id := c.Param("id")

	// Inital object
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create object from struct
	var pos models.Post

	// Find to database with id
	initializers.DB.First(&pos, id)

	// Update data
	initializers.DB.Model(&pos).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Response data
	c.JSON(200, gin.H{
		"post": pos,
	})

}

func DeleteDataById(c *gin.Context) {

	// Set Parameter
	id := c.Param("id")

	// Delete data
	initializers.DB.Delete(&models.Post{}, id)

	// Response data
	c.Status(200)
}
