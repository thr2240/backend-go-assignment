package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/thr2240/backend-go-assignment/models"
)

func GetAllLeads(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var leads []models.Lead

	if err := db.Find(&leads).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, leads)
}

func GetLead(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var lead models.Lead

	if err := db.First(&lead, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, lead)
}

func CreateLead(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var lead models.Lead

	if err := c.ShouldBindJSON(&lead); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Create(&lead).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, lead)
}

func DeleteLead(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Delete(&models.Lead{}, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateLead(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var lead models.Lead

	if err := db.First(&lead, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&lead); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Save(&lead).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, lead)
}