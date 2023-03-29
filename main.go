package main

import (
	"courses-api/database"
	"courses-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	database.Migrate()

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
