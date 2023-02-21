package controllers

import (
	"fmt"
	"net/http"

	"github.com/HasnathJami/Food-Nutrition-Go-Fiber-Gorm-Postgres-Docker/models"
	"github.com/HasnathJami/Food-Nutrition-Go-Fiber-Gorm-Postgres-Docker/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

var validate = validator.New()

func AddNutrient(context *fiber.Ctx){
     nutrientModel := &models.Nutrient{}
	//var nutrientModel *models.Nutrient

	err := context.BodyParser(&nutrientModel)
	utils.CheckError(err, "Request Failed")

	//validationErr := validate.Struct(nutrientModel)
	//utils.CheckError(err, validationErr.Error())

	err = utils.GetDB().Create(nutrientModel).Error
	utils.CheckError(err, "Could Not Create Nutrient")

	utils.ProcessSuccessful(context, http.StatusOK, "Nutrient Has Been Added", nutrientModel)	
}

func GetAllNutrients(context *fiber.Ctx){
    nutrients := &[]models.Nutrient{}

	err := utils.GetDB().Find(nutrients).Error
	utils.CheckError(err, "Could Not Get Nutrients")

	utils.ProcessSuccessful(context, http.StatusOK, "Nutrients Have Fetched Successfully", nutrients)
}

func GetNutrientById(context *fiber.Ctx){
	id := context.Params("id")
    nutrient := &models.Nutrient{}

	utils.CheckEmptyId(context, id, http.StatusInternalServerError, "Id Can Not Be Empty")
	fmt.Println("the ID is", id)

	err := utils.GetDB().Where("id = ?", id).First(nutrient).Error

	utils.CheckError(err, "Could Not Get The Book")

    utils.ProcessSuccessful(context , http.StatusOK, "Nutrient Id Fetched Successfully", nutrient)

}

func UpdateNutrient(context *fiber.Ctx){
    nutrient := &models.Nutrient{}

	err := context.BodyParser(nutrient)
	utils.CheckErrors(context, err, http.StatusUnprocessableEntity, "Request Failed")

	//validationErr := validate.Struct(nutrient)
	//utils.CheckError(err, validationErr.Error())

	id := context.Params("id")
	utils.CheckEmptyId(context, id, http.StatusInternalServerError, "Id Cannot Be Empty")

	err = utils.GetDB().Where("id = ?", id).Updates(nutrient).Error
	utils.CheckError(err, "Could Not Delete Nutrients")

	utils.ProcessSuccessful(context, http.StatusOK, "Nutrient Has Been Updated", id)

}

func DeleteNutrient(context *fiber.Ctx){
	nutrient := &models.Nutrient{}

	id := context.Params("id")
	utils.CheckEmptyId(context, id, http.StatusInternalServerError, "Id Cannot Be Empty")

	err := utils.GetDB().Delete(nutrient, id).Error
	utils.CheckError(err, "Could Not Delete Nutrients ")

    utils.ProcessSuccessful(context, http.StatusOK, "Nutrient Has Been Deleted", id)
}