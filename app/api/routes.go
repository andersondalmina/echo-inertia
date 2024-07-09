package api

import (
	api_controllers "go-inertia/app/api/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	group := e.Group("/api")
	group.GET("/providers", api_controllers.ProvidersController)
	group.GET("/providers/:id", api_controllers.ProviderController)
}
