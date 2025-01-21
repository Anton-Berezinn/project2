package repository

import (
	rp "rwa/projectfile/internal/model"
	storage "rwa/projectfile/internal/storage/postgres"
)

// AddWrapper - функция обертка.
func AddWrapper(r storage.Reposit, u rp.DataUser) int {
	id := r.Add(u)
	return id
}
