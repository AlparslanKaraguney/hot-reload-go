package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func welcome(c *fiber.Ctx) error {

	secretkey := os.Getenv("SECRET_KEY")

	return c.SendString(secretkey)
}

func main() {

	app := fiber.New()

	app.Get("/", welcome)

	app.Listen(":8000")
}
