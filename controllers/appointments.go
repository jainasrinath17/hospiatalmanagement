package controllers

import (
    "hospitalmanagement/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Controller struct {
    DB *gorm.DB
}

func (ac *Controller) GetAppointments(c *gin.Context) {
    appointments, err := models.GetAllAppointments(ac.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, appointments)
}

func (ac *Controller) CreateAppointment(c *gin.Context) {
    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := appointment.CreateAppointment(ac.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, appointment)
}
