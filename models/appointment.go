package models

import (
    "gorm.io/gorm"
)

type Appointment struct {
    gorm.Model
    PatientID uint    `json:"patient_id"`
    DoctorID  uint    `json:"doctor_id"`
    Date      string  `json:"date"`
    Patient   Patient `json:"patient" gorm:"foreignKey:PatientID"`
    Doctor    Doctor  `json:"doctor" gorm:"foreignKey:DoctorID"`
}

func (a *Appointment) CreateAppointment(db *gorm.DB) error {
    return db.Create(a).Error
}

func GetAllAppointments(db *gorm.DB) ([]Appointment, error) {
    var appointments []Appointment
    err := db.Preload("Patient").Preload("Doctor").Find(&appointments).Error
    return appointments, err
}
