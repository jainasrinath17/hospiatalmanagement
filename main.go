package main

import (
    "hospitalmanagement/models"
    "hospitalmanagement/pkg"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    dsn := "username:password@tcp(127.0.0.1:3306)/hospitaldb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Migrate the models
    db.AutoMigrate(&models.Patient{}, &models.Doctor{}, &models.Appointment{})

    // Set up the router
    router := pkg.SetupRouter(db)

    // Run the server
    router.Run(":8080")
}
