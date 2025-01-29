package postgres_articles

import (
	Articl "rwa/internal/handlers/response_articles"
)

type Reposit struct {
	Articles      []Articl.Article `json:"articles"`
	ArticlesCount int              `json:"articlesCount"`
}

type I interface {
	Add(data Articl.Article) error
	FindAll() ([]Articl.Article, int)
}

func (r *Reposit) Add(data Articl.Article) error {
	r.Articles = append(r.Articles, data)
	r.ArticlesCount += 1
	return nil
}

func (r Reposit) FindAll() Reposit {
	return r
}

func (r *Reposit) GetAuthor(name string) Reposit {
	answer := Reposit{}
	for _, v := range r.Articles {
		if v.Author.Username == name {
			answer.Articles = append(answer.Articles, v)
			answer.ArticlesCount += 1
		}
	}
	return answer
}

func (r *Reposit) GetTag(tagName string) Reposit {
	answer := Reposit{}
	for _, v := range r.Articles {
		for _, value := range v.TagList {
			if value == tagName {
				answer.Articles = append(answer.Articles, v)
				answer.ArticlesCount += 1
				break
			}
		}
	}
	return answer
}
