package config

import (
	"fmt"
)

// ConfigNew - возвращает секретный ключ
func ConfigNew() (string, error) {

	secretKey := "key"
	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY is not set")
	}

	return secretKey, nil
}
