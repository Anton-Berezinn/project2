package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"rwa/internal/model"
)

type I interface {
	ReadBody(r io.ReadCloser) (model.DataUser, error)
	ReadBodyArticle(r io.ReadCloser, user model.DataUser) (Article, error)
}

type Request struct{}

// ReadBody - метод,для чтения данных пользователя в структуру User, отдает ответ вложенную структуру.
func (req *Request) ReadBody(r io.ReadCloser) (model.DataUser, error) {
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

// ReadBodyArticle - метод, для чтения данных по артиклу.
func (req *Request) ReadBodyArticle(r io.ReadCloser, user model.DataUser) (Article, error) {
	u := &Artic{}
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
