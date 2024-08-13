package models

import (
    "gorm.io/gorm"
)

type Patient struct {
    gorm.Model
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Address string `json:"address"`
    Contact string `json:"contact"`
}

// Create a new patient in the database
func (p *Patient) CreatePatient(db *gorm.DB) error {
    return db.Create(p).Error
}

// Get all patients from the database
func GetAllPatients(db *gorm.DB) ([]Patient, error) {
    var patients []Patient
    err := db.Find(&patients).Error
    return patients, err
}
