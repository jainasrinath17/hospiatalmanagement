package pkg

import (
	"hospitalmanagement/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	controllers := controllers.Controller{DB: db}

	// Patient routes
	router.GET("/patients", controllers.GetPatients)
	router.POST("/patients", controllers.CreatePatient)

	// Doctor routes
	router.GET("/doctors", controllers.GetDoctors)
	router.POST("/doctors", controllers.CreateDoctor)

	// Appointment routes
	router.GET("/appointments", controllers.GetAppointments)
	router.POST("/appointments", controllers.CreateAppointment)

	return router
}
