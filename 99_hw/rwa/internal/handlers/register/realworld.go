package register

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	rq "rwa/internal/handlers/request"
	resp "rwa/internal/handlers/response"
	storage "rwa/internal/storage/postgres"
	wrapper "rwa/internal/storage/repository"
	token "rwa/internal/token/jwt"
	"sync"
)

// Handler - для хранения jwt
type Handler struct {
	Data storage.Reposit
	Token
	//Token jwt.Token
}

type Token struct {
	Data  map[string]int
	Count int
	Mu    sync.Mutex
}

// Register - хэндлер принимать данные пользователя и отдает ответ
func (u *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := rq.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	id := wrapper.AddWrapper(u.Data, data)
	//Игнорируем ошибку, потому что всегда прийдет nil.
	_ = token.CreateToken(u.Token.Data, &u.Token.Count, id, u.Token.Mu)
	value, err := resp.AnswerUser(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	fmt.Println(u.Token.Data)
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
	id, b := token.CheckToken(h.Token.Data, header, h.Token.Mu)
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
	id, b := token.CheckToken(h.Token.Data, header, h.Token.Mu)
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

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
