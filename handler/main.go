package handler

import (
	"github.com/UnikMask/gofeedsite/view/mainpage"
	"github.com/labstack/echo/v4"
)

func HandleMainPageShow(c echo.Context) error {
	return mainpage.MainPage().Render(c.Request().Context(), c.Response())
}

func HandleLoginPageShow(c echo.Context) error {
	return mainpage.LoginPageShow().Render(c.Request().Context(), c.Response())
}
