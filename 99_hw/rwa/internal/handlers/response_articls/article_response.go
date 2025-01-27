package response_articls

import (
	"encoding/json"
	"rwa/internal/model"
	"time"
)

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u model.TestArticle, user model.DataUser) ([]byte, error) {
	answer := model.Artic{
		Article: model.TestArticle{
			Author: model.TestProfile{
				Username:  user.Username,
				Bio:       user.Bio,
				CreatedAt: time.Now(),
			},
			Body:        u.Body,
			Title:       u.Title,
			Description: u.Description,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
			TagList:     u.TagList,
		},
	}
	data, err := json.Marshal(answer)
	if err != nil {
		return nil, err
	}
	return data, nil
}
