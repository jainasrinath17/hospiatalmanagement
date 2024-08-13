package models

import (
	"hospitalmanagement/utils"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"type:varchar(20);unique"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

func (d *Doctor) CreateDoctor(db *gorm.DB) error {
	d.UUID, _ = utils.GenerateRandomUUID()
	return db.Create(d).Error
}

func GetAllDoctors(db *gorm.DB) ([]Doctor, error) {
	var doctors []Doctor
	err := db.Find(&doctors).Error
	return doctors, err
}

func GetDoctorByUUID(db *gorm.DB, uuid string) (Doctor, error) {
	var doctor Doctor
	err := db.First(&doctor, "uuid = ?", uuid).Error
	return doctor, err
}

func UpdateDoctorByUUID(db *gorm.DB, uuid string, updatedData Doctor) error {
	return db.Model(&Doctor{}).Where("uuid = ?", uuid).Updates(updatedData).Error
}

func DeleteDoctorByUUID(db *gorm.DB, uuid string) error {
	return db.Delete(&Doctor{}, "uuid = ?", uuid).Error
}
