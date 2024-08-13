package controllers

import (
	"hospitalmanagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *Controller) GetPatients(c *gin.Context) {
	patients, err := models.GetAllPatients(pc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (pc *Controller) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := patient.CreatePatient(pc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func (pc *Controller) GetPatientByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	patient, err := models.GetPatientByUUID(pc.DB, uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (pc *Controller) UpdatePatientByUUID(c *gin.Context) {
    uuid := c.Param("uuid")

    var updatedPatient models.Patient
    if err := c.ShouldBindJSON(&updatedPatient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.UpdatePatientByUUID(pc.DB, uuid, updatedPatient); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found or update failed"})
        return
    }

    c.JSON(http.StatusOK, updatedPatient)
}

func (pc *Controller) DeletePatientByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := models.DeletePatientByUUID(pc.DB, uuid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted Patient")
}
