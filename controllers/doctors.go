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
