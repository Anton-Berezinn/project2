package repository

import (
	rp "rwa/internal/model"
	storage "rwa/internal/storage/postgres"
)

// AddWrapper - функция обертка.
func AddWrapper(r storage.Reposit, u rp.DataUser) int {
	return r.Add(u)
}

// CheckWrapper- функция, обертка для проверки юзера
func CheckWrapper(r storage.Reposit, u rp.DataUser) (string, bool) {
	return r.Check(u)
}

func GetWrapper(r storage.Reposit, id int) (rp.DataUser, bool) {
	return r.GetCr(id)
}

func UpdateWrapper(r storage.Reposit, id int, user rp.DataUser) (rp.DataUser, bool) {
	return r.Update(id, user)
}
