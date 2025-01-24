package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/echo-rest-api/handlers"
)


func ResgisterRoutes(e *echo.Echo) {
	e.GET("/",handlers.HelloWorld)
	e.GET("/users",	handlers.GetUser)
	e.GET("/users/:id" , handlers.GetUserByID)

}


