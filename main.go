package main

import (
	"encoding/json"
	"fmt"
	"hospitalmanagement/models"
	"hospitalmanagement/pkg"
	"io"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	DBName    string `json:"dbname"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parseTime"`
	Loc       string `json:"loc"`
}

// LoadDBConfig loads the database configuration from a JSON file
func LoadDBConfig() (*DBConfig, error) {
	file, err := os.Open("db.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config DBConfig
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {

    config, err := LoadDBConfig()
    if err != nil {
        log.Fatal("Failed to load database configuration:", err)
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
    config.Username, config.Password, config.Host, config.Port,
    config.DBName, config.Charset, config.ParseTime, config.Loc)	
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// To intialize the tables
	db.AutoMigrate(&models.Patient{}, &models.Doctor{}, &models.Appointment{})

	router := pkg.SetupRouter(db)

	// Runing the server
	router.Run(":8080")
}
