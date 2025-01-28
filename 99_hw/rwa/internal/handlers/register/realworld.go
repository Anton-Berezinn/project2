package register

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	rq "rwa/internal/handlers/request"
	article "rwa/internal/handlers/request_articls"
	resp "rwa/internal/handlers/response"
	art_resp "rwa/internal/handlers/response_articls"
	reposit "rwa/internal/handlers/services"
	token "rwa/internal/token/jwt"
	"sync"
)

// Handler - для хранения jwt
type Handler struct {
	Repository reposit.UserService
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
	id := u.Repository.AddWrapper(data)
	//Игнорируем ошибку, потому что всегда прийдет nil.
	_ = token.CreateToken(u.Token.Data, &u.Token.Count, id, u.Token.Mu)
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
	name, ok := h.Repository.CheckWrapper(data)
	if !ok {
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
	id, ok := token.CheckToken(h.Token.Data, header, h.Token.Mu)
	if !ok {
		w.WriteHeader(401)
		return
	}
	user, ok := h.Repository.GetWrapper(id)
	if !ok {
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
	user, ok := h.Repository.UpdateWrapper(id, data)
	if !ok {
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

// CreateArticle- метод, для создания артикла
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(401)
		return
	}
	id, ok := token.CheckToken(h.Token.Data, header, h.Token.Mu)
	if !ok {
		w.WriteHeader(401)
		return
	}
	user, ok := h.Repository.GetWrapper(id)
	if !ok {
		w.WriteHeader(401)
		return
	}
	data, err := article.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	value, err := art_resp.AnswerUser(data, user)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write(value)

}
