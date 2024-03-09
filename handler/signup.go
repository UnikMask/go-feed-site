package handler

import (
	"log"

	"github.com/UnikMask/gofeedsite/databases"
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
		return err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Signup password hash generation error: %s", err)
		return err
	}

	err = databases.ExecutePreparedStatement("databases/sign_up.sql", u.Username, u.Email, u.FirstName, u.LastName, string(pass[:]))
	if err != nil {
		log.Printf("Signup prepared statement execution error: %s", err)
		return err
	}
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
