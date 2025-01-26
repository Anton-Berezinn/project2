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

// Get- метод для получения пользователя по индексу.
func (r *Reposit) Get(id int) (rp.DataUser, bool) {
	user, exists := r.DB[id]
	return *user, exists
}

// Update - метод для обновления данных пользователя.
func (r *Reposit) Update(id int, u rp.DataUser) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.DB[id]; exists {
		r.DB[id] = &u
		return true
	}
	return false
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
