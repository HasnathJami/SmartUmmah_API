package controllers

import (
	"net/http"

	"github.com/HasnathJami/Smart-Ummah/models"
	"github.com/HasnathJami/Smart-Ummah/utils"
	"github.com/gofiber/fiber"
)

//var validate = validator.New()

func GetAllHadis(context *fiber.Ctx){
    allHadis := &[]models.Hadis{}

	err := utils.GetDB().Find(allHadis).Error
	utils.CheckError(context, err, "Could Not Get All Hadis", http.StatusBadRequest)
	

	utils.ProcessSuccessful(context, http.StatusOK, "All Hadis Have Fetched Successfully", allHadis)
}