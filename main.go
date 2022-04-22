package main

import (
	"gorber/setup"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setup.Environment()
	setup.Database()
	setup.Routes(app)

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080"
	}
	app.Listen(":" + apiPort)
}
