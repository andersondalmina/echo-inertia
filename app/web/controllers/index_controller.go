package web_controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexController(c echo.Context) error {
	return c.Render(http.StatusOK, "Index", map[string]interface{}{
		"teamName": "GloboID - Família e operadoras!",
	})
}
