package handler

import (
	"fmt"
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserSignUp struct {
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
}

func HandleSignUp(c echo.Context) error {
	u := UserSignUp{}
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
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Signup password hash generation error: %s", err)
		return render(c, components.InputError([]string{
			"Internal Server Error Occured - please try again later."}))
	}

	err = databases.ExecutePreparedStatement("databases/sign_up.sql", u.Username, u.Email, u.FirstName, u.LastName, string(pass[:]))
	if err != nil {
		log.Printf("Signup prepared statement execution error: %s", err)
		return render(c, components.InputError([]string{
			"Internal Server Error Occured - please try again later."}))
	}
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
