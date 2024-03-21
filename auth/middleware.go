package auth

import (
	"context"
	"net/http"

	"github.com/UnikMask/gofeedsite/model"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ss, err := c.Cookie(USER_SESSION_NAME)
		if err != nil {
			return next(c)
		}
		u, ok := ValidateJwtToken(ss.Value)
		if ok {
			ctx := context.WithValue(c.Request().Context(), CTX_USER_AUTH, u)
			c.SetRequest(c.Request().WithContext(ctx))
		}
		return next(c)
	}
}

func StrictAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, ok := c.Request().Context().Value(CTX_USER_AUTH).(model.UserAuth)
		if !ok {
			c.Response().WriteHeader(http.StatusNoContent)
			return nil
		}
		return next(c)
	}
}
