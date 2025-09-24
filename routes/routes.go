package routes

import (
	"github.com/kukuh01/crud-golang/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/products", controllers.GetProducts)
	app.Get("/products/:id", controllers.GetProduct)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)
}
