package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/UnikMask/gofeedsite/posts"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/UnikMask/gofeedsite/view/layout"
	"github.com/labstack/echo/v4"
)

func AttachFeedHandlers(app *echo.Echo) {
	page := app.Group("/feed")
	page.Use(RedirectAuthPageMiddleware)
	page.GET("", HandleFeedPage)

	api := app.Group(model.ENDPOINT_FEED)
	api.Use(auth.StrictAuthMiddleware)
	api.GET("", HandleGetFeed)
}

func HandleFeedPage(c echo.Context) error {
	return render(c, layout.FeedPage())
}

func HandleGetFeed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		if err != nil {
			fmt.Printf("Error fetching page parameter: %v\n", err)
		}
		page = 1
	}
	id := auth.GetUserFromContextOrNone(c.Request().Context()).Id
	feed, err := posts.GetPosts(id, page)
	if err != nil {
		fmt.Printf("Failed to get user feed for user %d: %v", id, err)
		return render(c, layout.PostError("Failed to get posts - try again later!"))
	}
	return render(c, components.FeedSegment(page, feed))
}

func RedirectAuthPageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := c.Request().Context().Value(auth.CTX_USER_AUTH).(model.UserAuth)
		if !ok {
			c.Response().Header().Set("HX-Redirect", "/")
			c.Response().WriteHeader(http.StatusSeeOther)
			return render(c, layout.Redirection("http://localhost:3000/"))
		}
		return next(c)
	}
}
