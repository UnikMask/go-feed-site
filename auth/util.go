package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	ISSUER            string        = "go-basic-feed-site"
	USER_SESSION_NAME string        = "user_session"
	TOKEN_DURATION    time.Duration = 24 * time.Hour
)

type User struct {
	Id        int
	Username  string `form:"username"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
}

type BearerClaims struct {
	jwt.RegisteredClaims
}

type UserSession struct {
	Token     *jwt.Token
	ExpiresAt time.Time
}

func GetContextJwtToken(c context.Context) *jwt.Token {
	unparsedToken := c.Value(USER_SESSION_NAME).(string)
	token, err := jwt.ParseWithClaims(unparsedToken, &BearerClaims{}, keyFunc)
	if err != nil {
		log.Printf("Error Parsing JWT Token: %s", err.Error())
		return nil
	}
	return token
}

func CreateJwtToken(u User) UserSession {
	duration := time.Now().Add(TOKEN_DURATION)
	claims := BearerClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(duration),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    ISSUER,
			Subject:   u.Email,
			ID:        os.Getenv("HOST_ID"),
			Audience:  []string{os.Getenv("HOST_SITE")},
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
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
	cookie.Value = ss
	cookie.Expires = u.ExpiresAt
	c.SetCookie(cookie)
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("APP_PRIVATE_KEY")), nil
}
