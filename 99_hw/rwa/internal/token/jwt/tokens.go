package jwt

import (
	"strconv"
	"sync"
)

// CreateToken - функция, генерирует jwt токен.
//func CreateToke(id int, key string) (string, error) {
//	t := &Token{
//		Id: id,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
//		},
//	}
//	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, t).SignedString([]byte(key))
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}

// Для избежания data race
var mu sync.Mutex

// CreateToken - функция, для создания ключа.
func CreateToken(id int) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	if id == 0 {
		return "Token", nil
	}
	s := strconv.Itoa(id)
	return "Token" + s, nil
}

// CheckToken - функция, для проверки существования токена.
func CheckToken(m map[string]int, header string) (int, bool) {
	if _, ok := m[header]; !ok {
		return 0, false
	}
	return m[header], true
}
