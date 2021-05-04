package main

import (
	"github.com/kazuhe/bukuma-go/routes"
	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
