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
	passwordRegex *regexp.Regexp = compileRegex("[a-zA-Z0-9#?!@$ %^&*-;:,.\\(\\)\\[\\]]?")
	usernameRegex *regexp.Regexp = compileRegex("[a-zA-Z0-9_]?")
	verifiers     []Verification = []Verification{
		{"email", checkEmail, "Email Address is invalid!"},
		{"username", checkTooBig(32), "Username must be at most 32 characters!"},
		{"username", checkTooShort(6), "Username must be at least 6 characters!"},
		{"username", usernameRegex.MatchString, "Usernames can only contain digits, letters, and `_`!"},
		{"password", checkTooShort(8), "Password must not have less than 8 characters!"},
		{"password", checkTooBig(128), "Password must not have more than 128 characters!"},
		{"password", checkDigits, "Password must have at least 1 digit!"},
		{"password", checkLowers, "Password must have at least 1 lowercase letter!"},
		{"password", checkUppers, "Password must have at least 1 uppercase letter!"},
		{"password", checkSpecials, "Password have at least 1 special character!"},
		{"password", passwordRegex.MatchString, "Password must only contain letters, digits, and specials!"},
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

func checkEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func compileRegex(expr string) *regexp.Regexp {
	res, err := regexp.Compile(expr)
	if err != nil {
		log.Fatalf("Failed to compile regex: %s", err.Error())
	}
	return res
}

func checkTooBig(h int) func(string) bool {
	return func(value string) bool { return len(value) <= h }
}

func checkTooShort(l int) func(string) bool {
	return func(value string) bool { return len(value) >= l }
}

func checkUppers(value string) bool {
	for _, c := range value {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false

}

func checkLowers(value string) bool {
	for _, c := range value {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false

}

func checkDigits(value string) bool {
	for _, c := range value {
		if unicode.IsNumber(c) {
			return true
		}
	}
	return false

}

func checkSpecials(value string) bool {
	for _, c := range value {
		if unicode.IsPunct(c) || unicode.IsSymbol(c) {
			return true
		}
	}
	return false

}
