package web_controllers

import (
	"go-inertia/app/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProvidersController(c echo.Context) error {
	providers := repository.GetProviders()

	return c.Render(http.StatusOK, "Providers", map[string]interface{}{
		"providers": providers,
	})
}

func ProviderController(c echo.Context) error {
	id := c.Param("id")
	provider, err := repository.FindProvider(id)

	if err != nil {
		return c.Render(http.StatusOK, "Error", map[string]interface{}{"status": 404})
	}

	return c.Render(http.StatusOK, "Provider", map[string]interface{}{
		"provider": provider,
	})
}
