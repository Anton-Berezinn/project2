package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	config "rwa/projectfile/internal/config"
	rq "rwa/projectfile/internal/handlers/request"
	resp "rwa/projectfile/internal/handlers/response"
	storage "rwa/projectfile/internal/storage/postgres"
	wrapper "rwa/projectfile/internal/storage/repository"
	"rwa/projectfile/internal/token/jwt"
)

// Register - хэндлер принимать данные пользователя и отдает ответ
func (u *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := rq.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	id := wrapper.AddWrapper(u.Data, data)
	secretkey, err := config.ConfigNew()
	if err != nil {
		http.Error(w, "something went wrong in ConfigNew", http.StatusInternalServerError)
	}
	token, err := jwt.CreateToken(id, secretkey)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
		return
	}
	r.Header.Set("Authorization", token)
	value, err := resp.AnswerUser(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(value)
}

// Handler - для хранения jwt
type Handler struct {
	Data storage.Reposit
	jwt.Token
}

func GetApp() http.Handler {
	var h Handler
	storage := storage.NewMap()
	h.Data = *storage
	router := httprouter.New()
	router.POST("/api/users", h.Register)
	return router
}
