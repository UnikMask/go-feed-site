package handler

import (
	"net/http"
	"strconv"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/UnikMask/gofeedsite/posts"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/UnikMask/gofeedsite/view/layout"
	"github.com/labstack/echo/v4"
)

const (
	POST_ID_PARAM = "id"
)

func AttachPostHandlers(app *echo.Echo) {
	page := app.Group("/posts")
	page.GET("/:id", HandlePostPage)

	api := app.Group(model.ENDPOINT_POSTS)
    api.Use(auth.StrictAuthMiddleware)
	api.POST("/:id/like", HandleLikePost)
    api.GET("/:id/likes", HandleGetLikes)
}

func HandlePostPage(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
	}

	return render(c, layout.PostPage(postId))
}

func HandleLikePost(c echo.Context) error {
	post_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
	}
	user := auth.GetUserFromContextOrNone(c.Request().Context())
	post, status := posts.GetPost(post_id, user.Id)
	if status != posts.GET_POST_STATUS_OK {
		c.Response().WriteHeader(http.StatusNoContent)
		return nil
	}
	ok := posts.LikePost(post, user)
    if ok {
        post.Liked = !post.Liked
        if post.Liked {
            post.Likes += 1
        } else {
            post.Likes -= 1
        }
    }
	return render(c, components.LikeButton(post))
}

func HandleGetLikes(c echo.Context) error {
    post_id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.Response().WriteHeader(http.StatusBadRequest)
    }
    likes, ok := posts.GetLikes(post_id)
    if !ok {
        c.Response().WriteHeader(http.StatusNoContent)
        return nil
    }
    return render(c, components.LikeCounter(post_id, likes))
}
