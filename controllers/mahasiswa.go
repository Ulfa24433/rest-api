package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raufhm/rest-api/models"
	"gorm.io/gorm"
)

type MahasiswaInput struct {
	Nim  string `json:"nim"`
	Nama string `json:"nama"`
}

//tampil semua data

func MahasiswaTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"data": mhs})
}
