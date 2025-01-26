package request

import (
	"encoding/json"
	"fmt"
	"io"
	"rwa/internal/model"
)

// ReadBody - функция,для чтения данных пользователя в структуру User, отдает ответ вложенную структуру.
func ReadBody(r io.ReadCloser) (model.DataUser, error) {
	u := &model.Userr{}
	resp, err := io.ReadAll(r)
	if err != nil {
		return u.DataUser, fmt.Errorf("error in read %w", err)
	}
	err = json.Unmarshal(resp, u)
	if err != nil {
		return u.DataUser, fmt.Errorf("error in unmarshal %w", err)
	}
	return u.DataUser, nil
}
