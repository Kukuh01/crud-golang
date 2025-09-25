package main

import (
	// "html"
	"log"
	"os"

	"github.com/kukuh01/crud-golang/config"
	"github.com/kukuh01/crud-golang/database/migrations"
	"github.com/kukuh01/crud-golang/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	config.InitDatabase()
	migrations.RunMigration()

	//Load template
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// register routes
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	log.Fatal(app.Listen(":" + port))
}
