package handler

import (
	"github.com/labstack/echo/v4"
)

func AttachFormHandlers(app *echo.Echo) {
	endpoint := app.Group("api/forms")
	// Forms Verification
	endpoint.POST("/verify/username", attachFormsVerifyType("username", UsernameVerifiers))
	endpoint.POST("/verify/email", attachFormsVerifyType("email", EmailVerifiers))
	endpoint.POST("/verify/password", attachFormsVerifyType("password", PasswordVerifiers))
	endpoint.POST("/verify/firstname", attachFormsVerifyType("firstname", NameVerifiers))
	endpoint.POST("/verify/lastname", attachFormsVerifyType("lastname", NameVerifiers))

	// Login-Signup
	endpoint.POST("/signup", HandleSignUp)
	endpoint.POST("/login", HandleLogIn)
}
