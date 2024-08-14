package pkg

import (
	"hospitalmanagement/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	controllers := controllers.Controller{DB: db}

	p := router.Group("/patient")
	{
		p.GET("/", controllers.GetPatients)
		p.POST("/", controllers.CreatePatient)
		p.GET("/:uuid", controllers.GetPatientByUUID)
		p.PATCH("/:uuid", controllers.UpdatePatientByUUID)
		p.DELETE("/:uuid", controllers.DeletePatientByUUID)
	}

	d := router.Group("/doctor")
	{
		d.GET("/", controllers.GetDoctors)
		d.POST("/", controllers.CreateDoctor)
		d.GET("/:uuid", controllers.GetDoctorByUUID)
		d.PATCH("/:uuid", controllers.UpdateDoctorByUUID)
		d.DELETE("/:uuid", controllers.DeleteDoctorByUUID)
	}

	a := router.Group("/appointments")
	{
		a.GET("/", controllers.GetAppointments)
		a.PATCH("/:uuid", controllers.UpdateAppointmentByUUID)
		a.POST("/", controllers.CreateAppointment)
		a.DELETE("/:uuid", controllers.DeleteAppointmentByUUID)
	}

	return router
}
