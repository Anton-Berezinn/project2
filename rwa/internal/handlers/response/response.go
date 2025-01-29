package response

import (
	"encoding/json"
	"rwa/internal/model"
	"time"
)

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u model.DataUser, token string) ([]byte, error) {
	answer := model.Response{
		User: model.TestProfile{
			ID:        u.ID,
			Email:     u.Email,
			Username:  u.Username,
			CreatedAt: time.Now(),
			Bio:       u.Bio,
			Token:     token,
		},
	}
	data, err := json.Marshal(answer)
	if err != nil {
		return nil, err
	}
	return data, nil
}
