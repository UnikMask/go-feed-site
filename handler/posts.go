package handler

import (
	"net/http"
	"strconv"

	"github.com/UnikMask/gofeedsite/view/layout"
	"github.com/labstack/echo/v4"
)

const (
    POST_ID_PARAM = "id"
)

func AttachPostHandlers(app *echo.Echo) {
    endpoints := app.Group("/posts")
    endpoints.GET("/:id" , HandlePostPage)
}

func HandlePostPage(c echo.Context) error {
    postId, err := strconv.Atoi(c.Param("id")) 
    if err != nil {
        c.Response().WriteHeader(http.StatusBadRequest)
    }

    return render(c, layout.PostPage(postId));
}
