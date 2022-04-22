package actor

import (
	"gorber/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/actor", getActors)
	app.Get("/actor/:id", getActor)
	app.Post("/actor", newActor)
	app.Delete("/actor/:id", deleteActor)
}

func getActors(c *fiber.Ctx) error {
	db := database.Conn
	var actors []database.Actor
	db.Find(&actors)
	return c.JSON(actors)
}

func getActor(c *fiber.Ctx) error {
	db := database.Conn
	var actor database.Actor
	db.First(&actor, c.Params("id"))
	return c.JSON(actor)
}

func newActor(c *fiber.Ctx) error {
	db := database.Conn
	var actor database.Actor
	if err := c.BodyParser(&actor); err != nil {
		return err
	}
	db.Create(&actor)
	return c.JSON(actor)
}

func deleteActor(c *fiber.Ctx) error {
	db := database.Conn
	var actor database.Actor
	db.First(&actor, c.Params("id"))
	db.Delete(&actor)
	return c.JSON(actor)
}
