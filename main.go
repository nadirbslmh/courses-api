package main

import (
	"courses-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // buat instance aplikasi echo baru

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323")) // jalankan server
}
