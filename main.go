package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raufhm/rest-api/controllers"
	"github.com/raufhm/rest-api/models"
)

func main() {

	r := gin.Default()

	// Model
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"Data": "Universitas Muhammadiyah Malang"})
	})

	r.GET("/mahasiswa", controllers.MahasiswaTampil)

	r.Run()
}
