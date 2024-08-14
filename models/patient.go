package models

import (
	"hospitalmanagement/utils"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	UUID    string `json:"uuid" gorm:"type:varchar(20);unique"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Contact string `json:"contact"`
}

func (p *Patient) CreatePatient(db *gorm.DB) error {
	p.UUID, _ = utils.GenerateRandomUUID()
	return db.Create(p).Error
}

func GetAllPatients(db *gorm.DB, limit, offset int) ([]Patient, error) {
    var patients []Patient
    err := db.Limit(limit).Offset(offset).Find(&patients).Error
    return patients, err
}

func GetPatientByUUID(db *gorm.DB, uuid string) (Patient, error) {
	var patient Patient
	err := db.First(&patient, "uuid = ?", uuid).Error
	return patient, err
}

func UpdatePatientByUUID(db *gorm.DB, uuid string, updatedData Patient) error {
	return db.Model(&Patient{}).Where("uuid = ?", uuid).Updates(updatedData).Error
}

func DeletePatientByUUID(db *gorm.DB, uuid string) error {
	return db.Delete(&Patient{}, "uuid = ?", uuid).Error
}
