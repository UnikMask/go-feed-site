package auth

import (
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(u User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Signup password hash generation error: %s", err)
		return err
	}

	err = databases.ExecutePreparedStatement("databases.sign_up.sql", u.Username, u.Email, u.FirstName, u.LastName, string(pass[:]))
	if err != nil {
		return err
	}
	return nil
}
