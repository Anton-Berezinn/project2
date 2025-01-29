package response_articles

import (
	"encoding/json"
	"rwa/internal/model"
	"time"
)

// AnswerUser - функция, для ответа юзеру
type Artic struct {
	Article Article `json:"article"`
}

type Article struct {
	Author         model.DataUser `json:"author"`
	Body           string         `json:"body"`
	CreatedAt      time.Time      `json:"createdAt"`
	Description    string         `json:"description"`
	Favorited      bool           `json:"favorited"`
	FavoritesCount int            `json:"favoritesCount"`
	Slug           string         `json:"slug" testdiff:"ignore"`
	TagList        []string       `json:"tagList"`
	Title          string         `json:"title"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u Article) ([]byte, error) {
	answer := Artic{
		Article: Article{
			Author: model.DataUser{
				Username: u.Author.Username,
				Bio:      u.Author.Bio,
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
