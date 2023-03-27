package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // buat instance aplikasi echo baru

	e.Logger.Fatal(e.Start(":1323")) // jalankan server
}
