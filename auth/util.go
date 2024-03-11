package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/UnikMask/gofeedsite/databases"
	"github.com/UnikMask/gofeedsite/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	ISSUER            string        = "go-basic-feed-site"
	USER_SESSION_NAME string        = "user_session"
	TOKEN_DURATION    time.Duration = 24 * time.Hour
	CTX_USER_AUTH     string        = "UserAuth"
)

type UserForm struct {
	Id        int
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
}

type UserAuth struct {
	Email string `json:"email"`
}

type BearerClaims struct {
	UserAuth
	jwt.RegisteredClaims
}

func (u UserAuth) GetUserAuth() UserAuth {
	return u
}

func (u UserForm) GetUserAuth() UserAuth {
	return UserAuth{Email: u.Email}
}

func (u UserAuth) GetUser() (model.User, bool) {
	res := model.User{}
	found, err := databases.QueryRow("databases/fetch_user_info.sql",
		[]any{u.Email},
		[]any{&res.Username, &res.Email, &res.FirstName, &res.LastName})
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return model.User{}, false
	}
	if !found {
		return model.User{}, false
	}
	return res, true
}

type UserSession struct {
	Token     *jwt.Token
	ExpiresAt time.Time
}

func ValidateJwtToken(ss string) (UserAuth, bool) {
	token, err := jwt.ParseWithClaims(ss, &BearerClaims{}, keyFunc)
	if err != nil {
		return UserAuth{}, false
	}
	claims, ok := token.Claims.(*BearerClaims)
	if !ok {
		return UserAuth{}, false
	}
	return claims.GetUserAuth(), true
}

func CreateJwtToken(u UserAuth) UserSession {
	duration := time.Now().Add(TOKEN_DURATION)
	claims := BearerClaims{
		UserAuth: u,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(duration),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    ISSUER,
			Subject:   u.Email,
			ID:        os.Getenv("HOST_ID"),
			Audience:  []string{os.Getenv("HOST_SITE")},
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return UserSession{token, duration}
}

func SignJwtToken(t *jwt.Token) (string, error) {
	ss, err := t.SignedString([]byte(os.Getenv("APP_PRIVATE_KEY")))
	if err != nil {
		log.Printf("Error signing JWT token: %s", err.Error())
		return "", err
	}
	return ss, nil
}

func SetAuthCookie(c echo.Context, u UserSession) {
	ss, err := u.Token.SignedString([]byte(os.Getenv("APP_PRIVATE_KEY")))
	if err != nil {
		log.Printf("Error signing JWT token: %s", err.Error())
		return
	}
	cookie := new(http.Cookie)
	cookie.Name = USER_SESSION_NAME
	cookie.Expires = u.ExpiresAt
	cookie.Value = ss
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("APP_PRIVATE_KEY")), nil
}
