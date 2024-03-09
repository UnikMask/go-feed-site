package handler

import (
	"fmt"
	"log"

	"github.com/UnikMask/gofeedsite/auth"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/labstack/echo/v4"
)

func HandleSignUp(c echo.Context) error {
	u := auth.User{}
	err := c.Bind(&u)
	if err != nil {
		log.Printf("Signup binding error: %s", err)
		return render(c, components.InputError([]string{
			"Failed to sign up - some informations are invalid."}))
	}

	// Verify Content
	messages := []string{}
	messages = append(messages, GetMessages(EmailVerifiers, u.Email)...)
	messages = append(messages, GetMessages(PasswordVerifiers, u.Password)...)
	messages = append(messages, GetMessages(UsernameVerifiers, u.Username)...)
	messages = append(messages, GetMessages(NameVerifiers, u.FirstName)...)
	messages = append(messages, GetMessages(NameVerifiers, u.LastName)...)
	if len(messages) != 0 {
		return render(c, components.InputError([]string{
			"You have " + fmt.Sprint(len(messages)) + " errors in your form!"}))
	}

	// Hash password
	err = auth.SignUp(u)
	if err != nil {
		return render(c, components.InputError([]string{
			"Internal Server Error Occured - please try again later."}))
	}
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
