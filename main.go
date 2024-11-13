package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG!!!")
	})

	app.Get("/name", func(c *fiber.Ctx) error {
		return c.SendString(appName)
	})

	log.Fatal(app.Listen(":8080"))
}
