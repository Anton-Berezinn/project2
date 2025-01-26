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
	router := httprouter.New()
	router.POST("/api/users", h.Register)
	return router
}
