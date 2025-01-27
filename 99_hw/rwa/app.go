package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rwa/internal/handlers/register"
	storage "rwa/internal/storage/postgres"
)

func GetApp() http.Handler {
	var h register.Handler
	storage := storage.NewMap()
	h.Data = *storage
	h.M = map[string]int{}
	router := httprouter.New()
	router.POST("/api/users", h.Register)
	router.POST("/api/users/login", h.Login)
	router.GET("/api/user", h.Main)
	router.PUT("/api/user", h.Update)
	router.POST("/api/articles")
	return router
}
