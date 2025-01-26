package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rwa/internal/config"
	rq "rwa/internal/handlers/request"
	resp "rwa/internal/handlers/response"
	storage "rwa/internal/storage/postgres"
	wrapper "rwa/internal/storage/repository"
	"rwa/internal/token/jwt"
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
	w.WriteHeader(201)
	w.Write(value)
}

// Handler - для хранения jwt
type Handler struct {
	Data storage.Reposit
	jwt.Token
}
