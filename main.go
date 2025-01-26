package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/echo-rest-api/routes"
)

func main() {
	e := echo.New()
	//call routes
	routes.ResgisterRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}