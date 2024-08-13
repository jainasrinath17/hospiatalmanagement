package models

import (
	"hospitalmanagement/utils"

	"gorm.io/gorm"
)

type Appointment struct {
    gorm.Model
	UUID    string `json:"uuid" gorm:"type:varchar(20);unique"`
    PatientID uint    `json:"patient_id"`
    DoctorID  uint    `json:"doctor_id"`
    Date      string  `json:"date"`
    Patient   Patient `json:"patient" gorm:"foreignKey:PatientID"`
    Doctor    Doctor  `json:"doctor" gorm:"foreignKey:DoctorID"`
}

func (a *Appointment) CreateAppointment(db *gorm.DB) error {
    a.UUID, _ = utils.GenerateRandomUUID()
    return db.Create(a).Error
}

func GetAppointmentByUUID(db *gorm.DB, uuid string) (Appointment, error) {
    var appointment Appointment
    err := db.First(&appointment, "uuid = ?", uuid).Error
    return appointment, err
}

func GetAllAppointments(db *gorm.DB) ([]Appointment, error) {
    var appointments []Appointment
    err := db.Preload("Patient").Preload("Doctor").Find(&appointments).Error
    return appointments, err
}

func UpdateAppointmentByUUID(db *gorm.DB, uuid string, updatedData Appointment) error {
    return db.Model(&Appointment{}).Where("uuid = ?", uuid).Updates(updatedData).Error
}

func DeleteAppointmentByUUID(db *gorm.DB, uuid string) error {
	return db.Delete(&Appointment{}, "uuid = ?", uuid).Error
}
