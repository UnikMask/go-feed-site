package auth

import (
	"context"
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
			ctx := context.WithValue(c.Request().Context(), "UserAuth", u)
			c.SetRequest(c.Request().WithContext(ctx))
		}

		return next(c)
	}
}
