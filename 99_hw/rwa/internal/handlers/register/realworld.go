package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	rq "rwa/internal/handlers/request"
	resp "rwa/internal/handlers/response"
	storage "rwa/internal/storage/postgres"
	wrapper "rwa/internal/storage/repository"
	token "rwa/internal/token/jwt"
)

// Handler - для хранения jwt
type Handler struct {
	Data storage.Reposit
	M    map[string]int
	//Token jwt.Token
}

// Register - хэндлер принимать данные пользователя и отдает ответ
func (u *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := rq.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	id := wrapper.AddWrapper(u.Data, data)
	token, err := token.CreateToken(id)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
		return
	}
	u.M[token] = id
	value, err := resp.AnswerUser(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Write(value)
}

// Login - метод для проверки аунтификации.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := rq.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	name, b := wrapper.CheckWrapper(h.Data, data)
	if !b {
		w.WriteHeader(404)
		return
	}
	data.Username = name
	value, err := resp.AnswerUser(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(value)

}

// Main- метод, для полученния данных пользователя.
func (h *Handler) Main(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(401)
		return
	}
	id, b := token.CheckToken(h.M, header)
	if !b {
		w.WriteHeader(401)
		return
	}
	user, b := wrapper.GetWrapper(h.Data, id)
	if !b {
		w.WriteHeader(401)
		return
	}
	value, err := resp.AnswerUser(user)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(value)

}

// Update - метод, для обновления данных пользователя.
func (h *Handler) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(401)
		return
	}
	id, b := token.CheckToken(h.M, header)
	if !b {
		w.WriteHeader(401)
		return
	}
	data, err := rq.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	user, b := wrapper.UpdateWrapper(h.Data, id, data)
	if !b {
		w.WriteHeader(401)
		return
	}
	value, err := resp.AnswerUser(user)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(value)

}
