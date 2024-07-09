package api_controllers

import (
	"go-inertia/app/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProvidersController(c echo.Context) error {
	providers := repository.GetProviders()

	return c.JSON(http.StatusOK, providers)
}

func ProviderController(c echo.Context) error {
	id := c.Param("id")
	provider, err := repository.FindProvider(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, provider)
}
