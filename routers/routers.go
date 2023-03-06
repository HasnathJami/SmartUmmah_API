package routers

import (
	"github.com/HasnathJami/Smart-Ummah/controllers"
	"github.com/gofiber/fiber"
)

func Router() *fiber.App {
   app := fiber.New()

   api := app.Group("/api")
   
   api.Get("/allHadis", controllers.GetAllHadis)

   return app
}