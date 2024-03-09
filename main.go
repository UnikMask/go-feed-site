package main

import (
	"context"
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	err := databases.OpenDatabase("app.db")
	if err != nil {
		log.Fatalf("Failed to open database: %s", err.Error())
	}
	err = databases.ExecuteFile("databases/on_startup.sql")
	if err != nil {
		databases.CloseDatabase()
		log.Fatalf("Failed to start up database: %s", err.Error())
	}

	app.Use(withUser)
	app.Static("/assets", "assets")
	app.GET("/", handler.HandleMainPageShow)
	app.GET("/login", handler.HandleLoginPageShow)
	app.GET("/user", userHandler.HandleUserShow)
	handler.AttachVerifyHandlers(app)
	app.Start(":3000")

	databases.CloseDatabase()
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "invalid@outlook.gg")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
