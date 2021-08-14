package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/usename-Poezd/mysite-api/internal/config"
	"github.com/usename-Poezd/mysite-api/internal/domain"
	"time"
)

func MakeToken(u *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_name"] = u.Name
	claims["user_id"] = u.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	conf, _ := config.GetConfig(".")
	t, err := token.SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return "nil", errors.New("token was not made")
	}

	return t, nil
}