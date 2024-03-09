package handler

import (
	"github.com/labstack/echo/v4"
)

func AttachFormHandlers(app *echo.Echo) {
	// Forms Verification
	app.POST("/forms/verify/username", attachFormsVerifyType("username", UsernameVerifiers))
	app.POST("/forms/verify/email", attachFormsVerifyType("email", EmailVerifiers))
	app.POST("/forms/verify/password", attachFormsVerifyType("password", PasswordVerifiers))
	app.POST("/forms/verify/firstname", attachFormsVerifyType("firstname", NameVerifiers))
	app.POST("/forms/verify/lastname", attachFormsVerifyType("lastname", NameVerifiers))

	// Login-Signup
	app.POST("/forms/signup", HandleSignUp)
}
