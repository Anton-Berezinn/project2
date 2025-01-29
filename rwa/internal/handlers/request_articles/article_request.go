package request_articls

import (
	"encoding/json"
	"fmt"
	"io"
	resp "rwa/internal/handlers/response_articles"
	"rwa/internal/model"
)

func ReadBody(r io.ReadCloser, user model.DataUser) (resp.Article, error) {
	u := &resp.Artic{}
	resp, err := io.ReadAll(r)
	if err != nil {
		return u.Article, fmt.Errorf("error in read %w", err)
	}
	err = json.Unmarshal(resp, u)
	if err != nil {
		return u.Article, fmt.Errorf("error in unmarshal %w", err)
	}
	u.Article.Author = model.DataUser{
		Username: user.Username,
		Bio:      user.Bio,
	}
	return u.Article, nil
}
