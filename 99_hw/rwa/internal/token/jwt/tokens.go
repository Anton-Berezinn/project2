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

// CreateToken - функция, для создания ключа.
func CreateToken(m map[string]int, count *int, id int, mu sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()
	if *count == 0 {
		m["Token"] = id
		*count++
		return nil
	}
	s := strconv.Itoa(*count + 1)
	m["token"+s] = id
	*count++
	return nil
}

// CheckToken - функция, для проверки существования токена.
func CheckToken(m map[string]int, header string, mu sync.Mutex) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := m[header]; !ok {
		return 0, false
	}
	return m[header], true
}
