package web

import (
	web_controllers "go-inertia/app/web/controllers"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	group := e.Group("")

	group.GET("/", web_controllers.IndexController)
	group.GET("/user", web_controllers.UserController)
	group.GET("/providers", web_controllers.ProvidersController)
	group.GET("/providers/:id", web_controllers.ProviderController)
}
