package setup

import (
	"fmt"
	"gorber/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func Environment() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Routes(app *fiber.App) {
	fmt.Printf("init routes")
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func Database() {
	fmt.Printf("init database")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbString := dbUser + ":" + dbPass + "@(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	database.Conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dbString,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %s as %s user. Listening on port %s.\n", dbName, dbUser, dbPort)

	database.Conn.Set("gorm:table_options", "ENGINE=InnoDB")
	database.Conn.AutoMigrate(
		&database.Movie{},
		&database.ProductionCompany{},
		&database.Actor{},
	)
}
