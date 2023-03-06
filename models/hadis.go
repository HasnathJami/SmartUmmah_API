package models

import "gorm.io/gorm"



type Hadis struct {
	gorm.Model
	ID             uint    `gorm:"primary key";autoIncrement"`
	Name           string `json:"name"`
	Description    string `json:"description"`

}

func MigrateDatabase(db *gorm.DB) (error) {
	err := db.AutoMigrate(&Hadis{})
	return err
}