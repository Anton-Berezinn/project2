package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rwa/internal/config"
	handler "rwa/internal/handlers/register"
	repository "rwa/internal/handlers/services"
	repositoryArticle "rwa/internal/handlers/services_articles"
	storage "rwa/internal/repository/postgres"
)

func GetApp() http.Handler {
	var h handler.Handler
	storage := storage.NewMap()
	h.Repository = repository.NewUserService(storage)
	h.Token.Data = map[string]int{}
	key, err := config.ConfigNew()
	if err != nil {
		//log
		panic(err)
	}
	h.SecretKey = key
	h.RepositoryArticle = repositoryArticle.NewUserServiceArticles()
	router := httprouter.New()
	router.POST("/api/users", h.Register)
	router.POST("/api/users/login", h.Login)
	router.GET("/api/user", h.Main)
	router.POST("/api/user/logout", h.Logout)
	router.PUT("/api/user", h.Update)
	router.POST("/api/articles", h.CreateArticle)
	router.GET("/api/articles/", h.GetArticles)
	return router
}
