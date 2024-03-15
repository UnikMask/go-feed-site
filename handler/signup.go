package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/labstack/echo/v4"
)

func HandleSignUp(c echo.Context) error {
	u := auth.UserForm{}
	err := c.Bind(&u)
	if err != nil {
		log.Printf("Signup binding error: %s", err)
		return render(c, components.InputError(
			"Some informations are invalid."))
	}

	// Verify Content
	messages := []string{}
	messages = append(messages, GetMessages(EmailVerifiers, u.Email)...)
	messages = append(messages, GetMessages(PasswordVerifiers, u.Password)...)
	messages = append(messages, GetMessages(UsernameVerifiers, u.Username)...)
	messages = append(messages, GetMessages(NameVerifiers, u.FirstName)...)
	messages = append(messages, GetMessages(NameVerifiers, u.LastName)...)
	if len(messages) != 0 {
		return render(c, components.InputError(
			"You have "+fmt.Sprint(len(messages))+" errors in your form!"))
	}

	// Hash password
    u_auth, err := auth.SignUp(u)
	if err != nil {
		log.Printf("Error during user signup: %s", err.Error())
		return render(c, components.InternalServerErrorTemplate)
	}
	auth.SetAuthCookie(c, auth.CreateJwtToken(u_auth))
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}

func HandleLogIn(c echo.Context) error {
	u := auth.UserForm{}
	err := c.Bind(&u)
	if err != nil {
		log.Printf("Login binding error: %v", err)
	}

	if len(u.Email) == 0 || len(u.Password) == 0 {
		return render(c, components.InputError("Please complete all fields."))
	}

	u_auth, ok, err := auth.LogIn(u.Email, u.Password)
	if err != nil {
		log.Printf("Error during user login: %s", err.Error())
		return render(c, components.InternalServerErrorTemplate)
	}
	if !ok {
		return render(c, components.InputError("Email or password incorrect!"))
	}
	auth.SetAuthCookie(c, auth.CreateJwtToken(u_auth))
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}

func HandleLogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Expires = time.Unix(0, 0)
	cookie.Name = auth.USER_SESSION_NAME
	cookie.Value = ""
	cookie.Path = "/"
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
