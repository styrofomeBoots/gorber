package movie

import (
	"gorber/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/movie", getMovies)
	app.Get("/movie/:id", getMovie)
	app.Post("/movie", newMovie)
	app.Delete("/movie/:id", deleteMovie)
}

func getMovies(c *fiber.Ctx) error {
	db := database.Conn
	var movies []database.Movie
	db.Find(&movies)
	return c.JSON(movies)
}

func getMovie(c *fiber.Ctx) error {
	db := database.Conn
	var movie database.Movie
	db.First(&movie, c.Params("id"))
	return c.JSON(movie)
}

func newMovie(c *fiber.Ctx) error {
	db := database.Conn
	var movie database.Movie
	if err := c.BodyParser(&movie); err != nil {
		return err
	}
	db.Create(&movie)
	return c.JSON(movie)
}

func deleteMovie(c *fiber.Ctx) error {
	db := database.Conn
	var movie database.Movie
	db.First(&movie, c.Params("id"))
	db.Delete(&movie)
	return c.JSON(movie)
}
