package main

import (
	"github.com/UnikMask/gofeedsite/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}
