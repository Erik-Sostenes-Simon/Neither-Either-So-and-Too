package routes

import (
	"github.com/Proyecto/proyecto_ingles/handlers"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	routes := e.Group("/v1/ingles")

	routes.GET("/", handlers.HomeHandler)
	routes.GET("/form", handlers.FormHandler)
	routes.POST("/answers", handlers.CreateHandler)
	routes.GET("/students", handlers.GetAllHandler)
}
