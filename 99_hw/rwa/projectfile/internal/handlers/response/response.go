package response

import (
	"encoding/json"
	"fmt"
	"rwa/projectfile/internal/model"
)

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u model.DataUser) ([]byte, error) {
	var answer model.Answer
	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("error in marshal %w", err)
	}
	err = json.Unmarshal(data, &answer)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshal %w", err)
	}
	data, err = json.Marshal(answer)

	return data, err
}
