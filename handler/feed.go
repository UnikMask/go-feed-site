package handler

import (
	"strconv"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/UnikMask/gofeedsite/view/layout"
	"github.com/labstack/echo/v4"
)

func AttachFeedHandlers(app *echo.Echo) {
	page := app.Group("/feed")
	page.GET("/", HandleFeedPage)

	api := app.Group(model.ENDPOINT_FEED)
	api.Use(auth.StrictAuthMiddleware)
}

func HandleFeedPage(c echo.Context) error {
	return render(c, layout.FeedPage())
}

func HandleGetFeed(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 0 {
		page = 0
	}
	return nil
}
