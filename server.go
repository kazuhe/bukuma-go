package main

import (
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
