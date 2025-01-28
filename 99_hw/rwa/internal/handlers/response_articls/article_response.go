package response_articls

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
	Author         Profile   `json:"author"`
	Body           string    `json:"body"`
	CreatedAt      time.Time `json:"createdAt"`
	Description    string    `json:"description"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Slug           string    `json:"slug" testdiff:"ignore"`
	TagList        []string  `json:"tagList"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type Profile struct {
	ID        string `json:"id" testdiff:"ignore"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Token     string `json:"token" testdiff:"ignore"`
	Following bool
}

// AnswerUser - функция, для ответа юзеру
func AnswerUser(u model.TestArticle, user model.DataUser) ([]byte, error) {
	answer := Artic{
		Article: Article{
			Author: Profile{
				Username: user.Username,
				Bio:      user.Bio,
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
