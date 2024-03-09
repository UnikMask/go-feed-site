package handler

import (
	"github.com/labstack/echo/v4"
)

func AttachFormHandlers(app *echo.Echo) {
	formTypes := []string{"username", "email", "password"}
	for _, formType := range formTypes {
		app.POST("/forms/verify/"+formType, attachFormsVerifyValue(formType))
	}
	app.POST("/forms/signup", HandleSignUp)
}
