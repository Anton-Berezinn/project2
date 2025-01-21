package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Token struct {
	Id int
	jwt.StandardClaims
}

// CreateToken - функция, генерирует jwt токен.
func CreateToken(id int, key string) (string, error) {
	t := &Token{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, t).SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}
