package web_controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserController(c echo.Context) error {
	return c.Render(http.StatusOK, "User", map[string]interface{}{
		"userName": "Ander",
	})
}
