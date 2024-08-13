package pkg

import (
	"hospitalmanagement/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	controllers := controllers.Controller{DB: db}

	apiV1 := router.Group("/patient")
	{
		apiV1.GET("/", controllers.GetPatients)
		apiV1.POST("/", controllers.CreatePatient)
		apiV1.GET("/:uuid", controllers.GetPatientByUUID)
		apiV1.PATCH("/:uuid", controllers.UpdatePatientByUUID)
		apiV1.DELETE("/:uuid", controllers.DeletePatientByUUID)
	}

	apiV2 := router.Group("/doctor")
	{
		apiV2.GET("/", controllers.GetDoctors)
		apiV2.POST("/", controllers.CreateDoctor)
		apiV2.GET("/:uuid", controllers.GetDoctorByUUID)
		apiV2.PATCH("/:uuid", controllers.UpdateDoctorByUUID)
		apiV2.DELETE("/:uuid", controllers.DeleteDoctorByUUID)
	}

	apiV3 := router.Group("/appointments")
	{
		apiV3.GET("/", controllers.GetAppointments)
		apiV3.PATCH("/:uuid", controllers.UpdateAppointmentByUUID)
		apiV3.POST("/", controllers.CreateAppointment)
	}

	return router
}
