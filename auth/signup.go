package auth

import (
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(u UserForm) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Signup password hash generation error: %s", err)
		return err
	}

	err = databases.ExecutePreparedStatement("databases/sign_up.sql", u.Username, u.Email, u.FirstName, u.LastName, string(pass[:]))
	if err != nil {
		return err
	}
	return nil
}

func LogIn(email string, password string) (bool, error) {
	var username string
	var password_hash string
	found, err := databases.QueryRow("databases/fetch_user.sql",
		[]any{email}, []any{&username, &password_hash})
	if err != nil {
		return false, err
	}
	if !found || bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)) != nil {
		return false, err
	}
	return true, nil
}
