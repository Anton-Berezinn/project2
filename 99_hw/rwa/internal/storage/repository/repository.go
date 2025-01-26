package repository

import (
	rp "rwa/internal/model"
	storage "rwa/internal/storage/postgres"
)

// AddWrapper - функция обертка.
func AddWrapper(r storage.Reposit, u rp.DataUser) int {
	return r.Add(u)
}
