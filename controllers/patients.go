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
