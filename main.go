package main

import (
	"context"

	"github.com/UnikMask/gofeedsite/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.Use(withUser)
	app.Static("/assets", "assets")
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "invalid@outlook.gg")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
