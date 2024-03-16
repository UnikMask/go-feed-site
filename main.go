package main

import (
	"log"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	err := databases.OpenDatabase("file:app.db?cache=shared&_fk=true")
	if err != nil {
		log.Fatalf("Failed to open database: %s", err.Error())
	}
	err = databases.ExecuteFile("databases/on_startup.sql")
	if err != nil {
		databases.CloseDatabase()
		log.Fatalf("Failed to start up database: %s", err.Error())
	}

	app.Use(auth.AuthMiddleware)
	app.Static("/assets", "assets")
	app.GET("/", handler.HandleMainPageShow)
	app.GET("/login", handler.HandleLoginPageShow)
	handler.AttachUserHandlers(app)
	handler.AttachFormHandlers(app)
	handler.AttachPostHandlers(app)
	app.Start(":3000")

	databases.CloseDatabase()
}
