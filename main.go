package main

import (
	"log"
	"os"

	"github.com/kukuh01/crud-golang/config"
	"github.com/kukuh01/crud-golang/database/migrations"
	"github.com/kukuh01/crud-golang/routes"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	// register routes
	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
