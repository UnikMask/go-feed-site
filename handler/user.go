package handler

import (
	"net/http"
	"strconv"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/UnikMask/gofeedsite/posts"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/UnikMask/gofeedsite/view/user"
	"github.com/labstack/echo/v4"
)

func AttachUserHandlers(app *echo.Echo) {
	endpoint := app.Group(model.ENDPOINT_USERS)
	endpoint.Use(auth.StrictAuthMiddleware)
	endpoint.GET("/actions", HandleUserActions)
	endpoint.GET("/hide-actions", HandleUserActionsHide)
	endpoint.POST("/logout", HandleLogOut)
    endpoint.POST("/:id/follow", HandleUserFollow)
    endpoint.POST("/:id/unfollow", HandleUserUnfollow)
}

func HandleUserActions(c echo.Context) error {
	return render(c, user.UserActions())
}

func HandleUserActionsHide(c echo.Context) error {
	return render(c, user.UserActionsHidden())
}

func HandleUserFollow(c echo.Context) error {
    return HandleUserFollowToggle(c, true)
}

func HandleUserUnfollow(c echo.Context) error {
    return HandleUserFollowToggle(c, false)
}

func HandleUserFollowToggle(c echo.Context, follow bool) error {
    followee_id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.Response().WriteHeader(http.StatusBadRequest)
        return nil
    }
    follower := auth.GetUserFromContextOrNone(c.Request().Context())
    ok := posts.FollowUserToggle(follower.Id, followee_id, !follow) 
    if !ok {
        c.Response().WriteHeader(http.StatusNoContent)
        return nil
    }
    return render(c, components.FollowButton(followee_id, follow))
}
