package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/UnikMask/gofeedsite/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

    err := openDatabase()
    if err != nil {
        log.Fatalf("Failed to open database: %s", err.Error())
    }
    err = executeFile("db/on_startup.sql", db)
    if err != nil {
        db.Close()
        log.Fatalf("Failed to start up database: %s", err.Error())
    }

	app.Use(withUser)
	app.Static("/assets", "assets")
	app.GET("/", handler.HandleMainPageShow)
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")

    db.Close()
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "invalid@outlook.gg")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func openDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		return err
	}
	return nil
}
