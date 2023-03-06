package routers

import (
	"github.com/HasnathJami/Smart-Ummah/controllers"
	"github.com/gofiber/fiber"
)

func Router() *fiber.App {
   app := fiber.New()

   api := app.Group("/api")
   //api.Post("/nutrient", controllers.AddNutrient)
   api.Get("/allHadis", controllers.GetAllHadis)
   // api.Get("/nutrient/:id", controllers.GetNutrientById)
   // api.Put("/nutrient/:id", controllers.UpdateNutrient)
   // api.Delete("/nutrient/:id", controllers.DeleteNutrient)

   return app
}