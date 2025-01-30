package services

import (
	"rwa/internal/dto"
	storage "rwa/internal/repository/postgres_articles"
)

type ArticleService struct {
	Storage storage.Reposit
}

// NewUserServiceArticles -создание.
func NewUserServiceArticles() ArticleService {
	return ArticleService{}
}

// AddWrapper- обертка.
func (a *ArticleService) AddWrappers(data dto.Article) error {
	return a.Storage.Add(data)
}

func (a *ArticleService) GetAllWrapper() storage.Reposit {
	return a.Storage.FindAll()
}

func (a *ArticleService) GetByNameWrapper(name string) storage.Reposit {
	return a.Storage.GetAuthor(name)
}

func (a *ArticleService) GetByTagWrapper(tagName string) storage.Reposit {
	return a.Storage.GetTag(tagName)
}
