package production

import (
	"gorber/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/production", getProductionCompanies)
	app.Get("/production/:id", getProductionCompany)
	app.Post("/production", newProductionCompany)
	app.Delete("/production/:id", deleteProductionCompany)
}

func getProductionCompanies(c *fiber.Ctx) error {
	db := database.Conn
	var productionCompanies []database.ProductionCompany
	db.Find(&productionCompanies)
	return c.JSON(productionCompanies)
}

func getProductionCompany(c *fiber.Ctx) error {
	db := database.Conn
	var productionCompany database.ProductionCompany
	db.First(&productionCompany, c.Params("id"))
	return c.JSON(productionCompany)
}

func newProductionCompany(c *fiber.Ctx) error {
	db := database.Conn
	var productionCompany database.ProductionCompany
	if err := c.BodyParser(&productionCompany); err != nil {
		return err
	}
	db.Create(&productionCompany)
	return c.JSON(productionCompany)
}

func deleteProductionCompany(c *fiber.Ctx) error {
	db := database.Conn
	var productionCompany database.ProductionCompany
	db.First(&productionCompany, c.Params("id"))
	db.Delete(&productionCompany)
	return c.JSON(productionCompany)
}
