package handler

import (
	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/UnikMask/gofeedsite/view/user"
	"github.com/labstack/echo/v4"
)

func AttachUserHandlers(app *echo.Echo) {
	endpoint := app.Group(model.ENDPOINT_USERS)
	endpoint.Use(auth.StrictAuthMiddleware)
	endpoint.GET("/actions", HandleUserActions)
	endpoint.GET("/hide-actions", HandleUserActionsHide)
	endpoint.POST("/logout", HandleLogOut)
}

func HandleUserActions(c echo.Context) error {
	return render(c, user.UserActions())
}

func HandleUserActionsHide(c echo.Context) error {
	return render(c, user.UserActionsHidden())
}
