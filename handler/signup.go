package handler

import (
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
	err := c.Bind(u)
	if err != nil {
		return err
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = databases.ExecutePreparedStatement("databases/sign_up.sql", u.Username, u.Email, string(pass))
	if err != nil {
		return err
	}
	c.Response().Header().Set("HX-Redirect", "/")
	return nil
}
