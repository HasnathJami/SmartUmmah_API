package models

import "gorm.io/gorm"

type Nutrient struct {
	ID             uint    `gorm:"primary key";autoIncrement"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	Class          *string `json:"class"`
	HealthBenefits *string `json:"healthbenefits"`
	//Source         Source         `json:"source"`
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&Nutrient{})
	return err
}