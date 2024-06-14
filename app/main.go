package main

import (
	"net/http"

	"github.com/elipzis/inertia-echo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "public/build")

	inertiaConfig := inertia.NewDefaultInertiaConfig(e)
	inertiaConfig.RootView = "app.html"
	inertiaConfig.TemplatesPath = "frontend/*.html"

	e.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Inertia: inertia.NewInertiaWithConfig(inertiaConfig),
	}))

	e.GET("/teste", func(c echo.Context) error {
		return c.Render(http.StatusOK, "Index", map[string]interface{}{
			"teste": "Teste 123",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
