package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ihsankarim/backend-brighted/config"
	"github.com/ihsankarim/backend-brighted/database"
	"github.com/ihsankarim/backend-brighted/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system env...")
	}
	_ = godotenv.Load()

	config.InitDB()
	app := fiber.New()
	database.Migrate()

	routes.Register(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
