package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

type Token struct {
	Id int
	jwt.StandardClaims
}

// CreateToken - функция, генерирует jwt токен.
func CreateToken(m map[string]int, id int, key string) (string, error) {
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
	m[token] = id
	return token, nil

}

func DecodeToken(tokenString string, key string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(t *jwt.Token) (interface{}, error) {
		// Проверка, что метод подписи правильный
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	// Если токен валиден, возвращаем данные
	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// CheckToken - функция, для проверки токена, если токен есть удаляем.
func CheckToken(m map[string]int, token string, mu sync.Mutex) (int, bool) {
	for k, id := range m {
		if k == token {
			return id, true
		}
	}
	return 0, false
}

// DeleteToken - функция, для удаления токена по id.
func DeleteTokenById(data map[string]int, id int) error {
	for token, v := range data {
		if v == id {
			delete(data, token)
		}
	}
	return nil
}
