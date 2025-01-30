package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	handler "rwa/internal/handlers/register"
)

func GetApp() http.Handler {
	h := handler.NewHandler()
	router := httprouter.New()
	router.POST("/api/users", h.Register)
	router.POST("/api/users/login", h.Login)
	router.GET("/api/user", h.Main)
	router.POST("/api/user/logout", h.Logout)
	router.PUT("/api/user", h.Update)
	router.POST("/api/articles", h.CreateArticle)
	router.GET("/api/articles", h.GetArticles)
	return router
}
