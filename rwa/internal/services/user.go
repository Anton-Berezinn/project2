package services

import (
	rp "rwa/internal/model"
	storage "rwa/internal/repository/postgres"
)

type UserService struct {
	Storage storage.Reposit
}

func NewUserService(r storage.Reposit) UserService {
	return UserService{Storage: r}
}

// AddWrapper - функция обертка.
func (u *UserService) AddWrapper(user rp.DataUser) int {
	return u.Storage.Add(user)
}

// CheckWrapper- функция, обертка для проверки юзера
func (u *UserService) CheckWrapper(user rp.DataUser) (string, int, bool) {
	return u.Storage.Check(user)
}

func (u *UserService) GetWrapper(id int) (rp.DataUser, bool) {
	return u.Storage.GetCr(id)
}

func (u *UserService) UpdateWrapper(id int, user rp.DataUser) (rp.DataUser, bool) {
	return u.Storage.Update(id, user)
}
