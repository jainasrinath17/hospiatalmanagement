package models

import (
    "gorm.io/gorm"
)

type Doctor struct {
    gorm.Model
    Name      string `json:"name"`
    Specialty string `json:"specialty"`
}

// Create a new doctor in the database
func (d *Doctor) CreateDoctor(db *gorm.DB) error {
    return db.Create(d).Error
}

// Get all doctors from the database
func GetAllDoctors(db *gorm.DB) ([]Doctor, error) {
    var doctors []Doctor
    err := db.Find(&doctors).Error
    return doctors, err
}
