package response

import (
	"encoding/json"
	"fmt"
	"rwa/internal/model"
)

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u model.DataUser) ([]byte, error) {
	var answer model.Response
	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("error in marshal %w", err)
	}
	err = json.Unmarshal(data, &answer.User)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshal %w", err)
	}
	answer.User.CreatedAt.Valid = true
	data, err = json.Marshal(answer.User)
	return data, err
}
