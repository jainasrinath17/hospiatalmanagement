package controllers

import (
	"hospitalmanagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (dc *Controller) GetDoctors(c *gin.Context) {
	doctors, err := models.GetAllDoctors(dc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}

func (dc *Controller) CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := doctor.CreateDoctor(dc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, doctor)
}

func (dc *Controller) GetDoctorByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	doctor, err := models.GetDoctorByUUID(dc.DB, uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, doctor)
}

func (dc *Controller) UpdateDoctorByUUID(c *gin.Context) {
	uuid := c.Param("uuid")

	var updatedDoctor models.Doctor
	if err := c.ShouldBindJSON(&updatedDoctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateDoctorByUUID(dc.DB, uuid, updatedDoctor); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, updatedDoctor)
}

func (dc *Controller) DeleteDoctorByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := models.DeleteDoctorByUUID(dc.DB, uuid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted Patient")
}
