package register

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	rq "rwa/internal/handlers/request"
	article "rwa/internal/handlers/request_articles"
	resp "rwa/internal/handlers/response"
	art_resp "rwa/internal/handlers/response_articles"
	reposit "rwa/internal/handlers/services"
	repositArt "rwa/internal/handlers/services_articles"
	token "rwa/internal/token/jwt"
	"strings"
	"sync"
)

// Todo: Куда положить Handler? я пытался положить в app.go, но потом импортнуть не мог.
type Handler struct {
	Repository        reposit.UserService
	RepositoryArticle repositArt.ArticleService
	SecretKey         string
	Token
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
	tokenname, err := token.CreateToken(u.Data, id, u.SecretKey)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
		return
	}

	value, err := resp.AnswerUser(data, tokenname)
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
	name, id, ok := h.Repository.CheckWrapper(data)
	if !ok {
		w.WriteHeader(401)
		return
	}
	//игнорируем ошибку, потому что всегда будет nil.
	_ = token.DeleteTokenById(h.Data, id)
	data.Username = name
	tokenName, err := token.CreateToken(h.Data, id, h.SecretKey)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
		return
	}
	value, err := resp.AnswerUser(data, tokenName)
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
	tokenName := strings.SplitN(header, " ", 2)
	id, ok := token.CheckToken(h.Token.Data, tokenName[1], h.Token.Mu)
	if !ok {
		w.WriteHeader(401)
		return
	}
	user, ok := h.Repository.GetWrapper(id)
	if !ok {
		w.WriteHeader(401)
		return
	}
	value, err := resp.AnswerUser(user, "")
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
	tokenName := strings.SplitN(header, " ", 2)
	id, b := token.CheckToken(h.Token.Data, tokenName[1], h.Token.Mu)
	if !b {
		w.WriteHeader(401)
		return
	}
	//игнорируем, всегда вернется nil.
	_ = token.DeleteTokenById(h.Data, id)
	tokenname, err := token.CreateToken(h.Data, id, h.SecretKey)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
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
	value, err := resp.AnswerUser(user, tokenname)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(value)
}

// // CreateArticle- метод, для создания артикла
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(401)
		return
	}
	tokenName := strings.SplitN(header, " ", 2)
	id, ok := token.CheckToken(h.Token.Data, tokenName[1], h.Token.Mu)
	if !ok {
		w.WriteHeader(401)
		return
	}
	user, ok := h.Repository.GetWrapper(id)
	if !ok {
		w.WriteHeader(500)
		return
	}
	data, err := article.ReadBody(r.Body, user)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}

	//Todo: положить данные
	//игнорируем ошибку, там всегда будет nil
	_ = h.RepositoryArticle.AddWrapper(data)
	value, err := art_resp.AnswerUser(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Write(value)

}

// GetArticles -handler.
func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Todo: я пытался изменить url, он все равно приходил сюда и если пробовать по p.ByName достать,будет пусто
	if r.URL.String() == "/api/articles/" {
		answer := h.RepositoryArticle.GetAllWrapper()
		data, err := json.Marshal(answer)
		if err != nil {
			http.Error(w, "something went wrong in GetArticles", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		w.Write(data)
		return
	} else {
		fmt.Println(r.URL.String())
		fmt.Println(r.Body, "here")
	}

}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := r.Header.Get("Authorization")
	if header == "" {
		w.WriteHeader(401)
		return
	}

	tokenName := strings.SplitN(header, " ", 2)
	id, ok := token.CheckToken(h.Token.Data, tokenName[1], h.Token.Mu)
	if !ok {
		w.WriteHeader(401)
		return
	}
	//игнорируем всегда будет nil
	_ = token.DeleteTokenById(h.Data, id)
	w.WriteHeader(200)
}
