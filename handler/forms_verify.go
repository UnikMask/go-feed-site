package handler

import (
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/labstack/echo/v4"
)

func AttachVerifyHandlers(app *echo.Echo) {
	app.POST("/forms/verify", HandleFormVerify)
}

func HandleFormVerify(c echo.Context) error {
	return components.InputError("updog!").Render(c.Request().Context(), c.Response())
}
