package config

import (
	"fmt"
	"os"
)

// ConfigNew - возвращает секретный ключ
func ConfigNew() (string, error) {

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY is not set")
	}

	return secretKey, nil
}
