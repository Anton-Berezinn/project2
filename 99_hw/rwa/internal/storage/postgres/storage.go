package postgres

import (
	rp "rwa/internal/model"
	"sync"
)

// Репозиторий для хранения пользователей
type Reposit struct {
	DB    map[int]*rp.DataUser // Карта для хранения данных пользователей по индексу
	Count int                  // Счетчик для индексации пользователей
	mu    sync.RWMutex         //мьютекс для избежания гонки
}

type I interface {
	Add(u rp.DataUser)
	Get(id int) (rp.DataUser, bool)
	Update(id int, u rp.DataUser) bool
	Delete(id int) bool
}

// NewMap- функция для создания нового репозитория.
func NewMap() *Reposit {
	return &Reposit{
		DB: make(map[int]*rp.DataUser),
	}
}

// Add- метод для добавления пользователя в репозиторий.
func (r *Reposit) Add(u rp.DataUser) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.DB[r.Count] = &u
	r.Count++
	return r.Count - 1
}

// GetCr- метод для получения пользователя по индексу.
func (r *Reposit) GetCr(id int) (rp.DataUser, bool) {
	user, exists := r.DB[id]
	if !exists {
		return rp.DataUser{}, false
	}
	a := rp.DataUser{
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
	}
	return a, true

}

func (r *Reposit) Check(user rp.DataUser) (string, bool) {
	for _, u := range r.DB {
		if u.Email == user.Email && u.Password == user.Password {
			return u.Username, true
		}
	}
	return "", false
}

// Update - метод для обновления данных пользователя.
func (r *Reposit) Update(id int, user rp.DataUser) (rp.DataUser, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if data, exists := r.DB[id]; exists {
		if user.Email != "" {
			data.Email = user.Email
		}
		if user.Bio != "" {
			data.Bio = user.Bio
		}
		u := rp.DataUser{
			Username: data.Username,
			Email:    data.Email,
			Bio:      data.Bio,
		}
		return u, true
	}
	return rp.DataUser{}, false
}

// Delete - метод для удаления пользователей.
func (r *Reposit) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.DB[id]; exists {
		delete(r.DB, id)
		return true
	}
	return false
}
