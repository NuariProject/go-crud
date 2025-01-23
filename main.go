package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {
	r := gin.Default()

	r.POST("/test", controllers.CreateData)
	r.PUT("/test/:id", controllers.UpdateDataById)
	r.DELETE("/test/:id", controllers.DeleteDataById)
	r.GET("/test", controllers.GetListData)
	r.GET("/test/:id", controllers.GetDataById)

	r.Run()
}
