package main

import (
	"go-inertia/app/api"
	"go-inertia/app/web"

	"github.com/elipzis/inertia-echo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "public/build")

	inertiaConfig := inertia.NewDefaultInertiaConfig(e)
	inertiaConfig.RootView = "index.html"
	inertiaConfig.TemplatesPath = "public/*.html"

	e.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Inertia: inertia.NewInertiaWithConfig(inertiaConfig),
	}))

	api.ConfigureRoutes(e)
	web.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
