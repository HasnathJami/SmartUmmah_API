package database

import (
	"fmt"

	"github.com/HasnathJami/Smart-Ummah/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *models.Config) (*gorm.DB, error) {
   dsn := fmt.Sprintf(
	"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",config.Host,config.Port,config.Password,config.User,config.DBName,config.SSLMode)
   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   //DisableForeignKeyConstraintWhenMigrating: true,

   if err != nil{
     return db, err
   }
   return db, nil

}