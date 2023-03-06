package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/HasnathJami/Smart-Ummah/database"
	"github.com/HasnathJami/Smart-Ummah/models"
	"github.com/HasnathJami/Smart-Ummah/routers"
	"github.com/HasnathJami/Smart-Ummah/utils"
)

func main() {
	err := godotenv.Load(".env")
	utils.CheckSimpleError(err,"")

	config := &models.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := database.NewConnection(config)
	utils.CheckSimpleError(err, "Could Not Load The Database")

	err = models.MigrateDatabase(db)
	utils.CheckSimpleError(err, "Could Not Migrate The Database")

	// var repository *models.Repository = &models.Repository{
	// 	DB : db,
	// }

	r := routers.Router()
	utils.SetDB(db)
	
	fmt.Println("Starting The Server on Port"+os.Getenv("PORT"))
	log.Fatal(r.Listen(":" + os.Getenv("PORT")))
}