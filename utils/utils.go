package utils

import (
	"log"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func CheckError(context *fiber.Ctx, err error, message string, status int) {
	if err != nil {
        if message != ""{
          context.Status(status).JSON(
			&fiber.Map{"message" : message,
		})
	
		} else{
          //log.Fatal(err)
		context.Status(status).JSON(
			&fiber.Map{"message" : err,
		})
		}	
	}else{
        
		context.Status(status).JSON(
			&fiber.Map{"message" : "Record not found",
		})

	}
}

func CheckSimpleError(err error, message string){
	if err != nil {
        if message != ""{
          log.Fatal(message)
		} else{
          log.Fatal(err)
		}
		}
	}
	

func CheckErrors(context *fiber.Ctx, err error, status int , message string){
	if err != nil {
		context.Status(status).JSON(
			&fiber.Map{"message" : message,
		})

		return
	}
}

func CheckEmptyId(context *fiber.Ctx, id string, status int, message string){
	if id == ""{
		context.Status(status).JSON(&fiber.Map{
			"message" : message,
		})

		return
	}
}

func ProcessSuccessful(context *fiber.Ctx, status int, message string, data any){
	context.Status(status).JSON(&fiber.Map{
		"message" : message,
		"data"    : data,
	})
}

var (	
	 database *gorm.DB
)

func SetDB(db *gorm.DB){
    database = db
}

func GetDB() *gorm.DB{
   return database
}