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
	Verifier     func(value string) bool
	ErrorMessage string
}

var (
	passwordRegex     *regexp.Regexp = compileRegex("^[a-zA-Z0-9#?!@$ %^&*-;:,\\.\\(\\)\\[\\]]*$")
	usernameRegex     *regexp.Regexp = compileRegex("^[a-zA-Z0-9_]*$")
	PasswordVerifiers                = []Verification{
		{checkTooShort(8), "Password must not have less than 8 characters!"},
		{checkTooBig(128), "Password must not have more than 128 characters!"},
		{checkDigits, "Password must have at least 1 digit!"},
		{checkLowers, "Password must have at least 1 lowercase letter!"},
		{checkUppers, "Password must have at least 1 uppercase letter!"},
		{checkSpecials, "Password have at least 1 special character!"},
		{passwordRegex.MatchString, "Password must only contain letters, digits, and specials!"},
	}
	EmailVerifiers    = []Verification{{checkEmail, "Email Address is invalid!"}}
	UsernameVerifiers = []Verification{
		{checkTooBig(32), "Username must be at most 32 characters!"},
		{checkTooShort(6), "Username must be at least 6 characters!"},
		{usernameRegex.MatchString, "Usernames can only contain digits, letters, and `_`!"},
	}
	NameVerifiers = []Verification{
		{usernameRegex.MatchString, "Names can only contain letters"}}
)

func GetMessages(verifiers []Verification, value string) []string {
	messages := []string{}
	for _, verif := range verifiers {
		if !verif.Verifier(value) {
			messages = append(messages, verif.ErrorMessage)
		}
	}
	return messages
}

func attachFormsVerifyType(typ string, verifier []Verification) func(c echo.Context) error {
	return func(c echo.Context) error {
		value := c.FormValue(typ)
		errorMessages := GetMessages(verifier, value)
		return components.InputErrors(errorMessages).Render(c.Request().Context(), c.Response())
	}
}

func HandleFormVerify(c echo.Context) error {
	return components.InputError("updog!").Render(c.Request().Context(), c.Response())
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
