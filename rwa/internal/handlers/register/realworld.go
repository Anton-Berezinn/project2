package register

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rwa/internal/config"
	"rwa/internal/dto"
	storageUser "rwa/internal/repository/postgres"
	storage "rwa/internal/repository/postgres_articles"
	"rwa/internal/services"
	token "rwa/internal/token/jwt"
	"strings"
	"sync"
)

type Handler struct {
	Repository        services.UserService
	RepositoryArticle services.ArticleService
	SecretKey         string
	Answer            dto.Answer
	Request           dto.Request
	*Token
}

type Token struct {
	Data  map[string]int
	Count int
	Mu    sync.Mutex
}

func NewTokenStorage() *Token {
	return &Token{
		Data: make(map[string]int),
	}
}

func NewHandler() *Handler {
	storage := storageUser.NewMap()
	return &Handler{
		SecretKey:         config.ConfigNew(),
		Token:             NewTokenStorage(),
		Repository:        services.NewUserService(storage),
		RepositoryArticle: services.NewUserServiceArticles(),
	}
}

// Register - хэндлер принимать данные пользователя и отдает ответ
func (h *Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := dto.ReadBody(r.Body)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	id := h.Repository.AddWrapper(data)
	tokenname, err := token.CreateToken(h.Data, id, h.SecretKey)
	if err != nil {
		http.Error(w, "something went wrong in CreateToken", http.StatusInternalServerError)
		return
	}

	value, err := dto.AnswerUser(data, tokenname)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Write(value)
}

// Login - метод для проверки аунтификации.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := dto.ReadBody(r.Body)
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
	value, err := dto.AnswerUser(data, tokenName)
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
	value, err := dto.AnswerUser(user, "")
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
	data, err := dto.ReadBody(r.Body)

	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}
	user, ok := h.Repository.UpdateWrapper(id, data)
	if !ok {
		w.WriteHeader(401)
		return
	}
	value, err := dto.AnswerUser(user, tokenname)
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
	data, err := dto.ReadBodyArticle(r.Body, user)
	if err != nil {
		http.Error(w, "something went wrong in ReadBody", http.StatusInternalServerError)
		return
	}

	//игнорируем ошибку, там всегда будет nil
	_ = h.RepositoryArticle.AddWrappers(data)
	value, err := dto.NewAnswerTag(data)
	if err != nil {
		http.Error(w, "something went wrong in AnswerUser", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Write(value)

}

// GetArticles -handler.
func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.URL.Query().Get("author")
	tag := r.URL.Query().Get("tag")
	data := storage.Reposit{}
	if tag != "" && name == "" {
		data = h.GetByTag(tag)
	}
	if name == "" && tag == "" {
		data = h.GetAll()
	}
	if name != "" && tag == "" {
		data = h.GetByAuthor(name)
	}
	value, err := dto.AnswerTag(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving articles: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(value)
}

func (h *Handler) GetAll() storage.Reposit {
	answer := h.RepositoryArticle.GetAllWrapper()
	return answer
}

func (h *Handler) GetByAuthor(name string) storage.Reposit {
	answer := h.RepositoryArticle.GetByNameWrapper(name)
	return answer
}

func (h *Handler) GetByTag(tag string) storage.Reposit {
	answer := h.RepositoryArticle.GetByTagWrapper(tag)
	return answer
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
