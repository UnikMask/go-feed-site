package handler

import (
	"log"
	"net/mail"
	"regexp"
	"unicode"

	"github.com/UnikMask/gofeedsite/view/components"
	"github.com/labstack/echo/v4"
)

type Verification struct {
	Name         string
	Verifier     func(value string) bool
	ErrorMessage string
}

var (
	passwordRegex *regexp.Regexp = func(v string) *regexp.Regexp {
		res, err := regexp.Compile(v)
		if err != nil {
			log.Fatalf("Failed to compile regex: %s", err.Error())
		}
		return res
	}("[a-zA-Z0-9#?!@$ %^&*-;:,.\\(\\)\\[\\]]")

	verifiers []Verification = []Verification{{"email", func(value string) bool {
		_, err := mail.ParseAddress(value)
		return err == nil
	}, "Email Address is invalid!"},
		{"username", func(value string) bool {
			return len(value) < 32
		}, "Username must be less than 32 characters!"},
		{"password", func(value string) bool {
			return len(value) >= 8
		}, "Password must not have less than 8 characters!"},
		{"password", func(value string) bool {
			return len(value) <= 128
		}, "Password must not have more than 128 characters!"},
		{"password", func(value string) bool {
			for _, c := range value {
				if unicode.IsNumber(c) {
					return true
				}
			}
			return false
		}, "Password must have at least 1 digit!"},
		{"password", func(value string) bool {
			for _, c := range value {
				if unicode.IsLower(c) {
					return true
				}
			}
			return false
		}, "Password must have at least 1 lowercase letter!"},
		{"password", func(value string) bool {
			for _, c := range value {
				if unicode.IsUpper(c) {
					return true
				}
			}
			return false
		}, "Password must have at least 1 uppercase letter!"},
		{"password", func(value string) bool {
			for _, c := range value {
				if unicode.IsPunct(c) || unicode.IsSymbol(c) {
					return true
				}
			}
			return false
		}, "Password have at least 1 special character!"},
		{"password", func(value string) bool {
			return passwordRegex.MatchString(value)
		}, "Password must only contain letters, digits, and specials!"},
	}
)

func AttachVerifyHandlers(app *echo.Echo) {
	formTypes := []string{"username", "email", "password"}
	for _, formType := range formTypes {
		app.POST("/forms/verify/"+formType, attachFormsVerifyValue(formType))
	}
}

func attachFormsVerifyValue(typ string) func(c echo.Context) error {
	verifications := []Verification{}
	for _, verif := range verifiers {
		if verif.Name == typ {
			verifications = append(verifications, verif)
		}
	}
	return func(c echo.Context) error {
		value := c.FormValue(typ)
		errorMessages := []string{}
		for _, verif := range verifications {
			if !verif.Verifier(value) {
				errorMessages = append(errorMessages, verif.ErrorMessage)
			}
		}
		return components.InputError(errorMessages).Render(c.Request().Context(), c.Response())
	}
}

func HandleFormVerify(c echo.Context) error {
	return components.InputError([]string{"updog!"}).Render(c.Request().Context(), c.Response())
}
