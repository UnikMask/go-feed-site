package auth

import (
	"log"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(u UserForm) (model.UserAuth, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Signup password hash generation error: %s", err)
		return model.UserAuth{}, err
	}

    u_auth := model.UserAuth{}
    info := []any{u.Email, u.Username, u.FirstName, u.LastName, string(pass[:])}
    _, err = databases.QueryRow("databases/sign_up.sql", info, []any{&u_auth.Id})
	if err != nil {
		return model.UserAuth{}, err
	}
	return u_auth, nil
}

func LogIn(email string, password string) (model.UserAuth, bool, error) {
	var password_hash string
    u_auth := model.UserAuth{}
	found, err := databases.QueryRow("databases/fetch_user_creds.sql",
		[]any{email}, []any{&u_auth.Id, &password_hash})
	if err != nil {
        return u_auth, false, err
	}
	if !found || bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)) != nil {
        return u_auth, false, err
	}
	return u_auth, true, nil
}
