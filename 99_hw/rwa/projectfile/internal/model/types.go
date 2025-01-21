package model

type User struct {
	DataUser `json:"user"`
}

// DataUser - структура, для хранения данных
type DataUser struct {
	ID       string `json:"id" testdiff:"ignore"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Time struct {
	Valid bool
}

// Answer - структура для ответа
type Answer struct {
	ID        string `json:"id" testdiff:"ignore"`
	Email     string `json:"email"`
	CreatedAt Time   `json:"createdAt"`
	UpdatedAt Time   `json:"updatedAt"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Token     string `json:"token" testdiff:"ignore"`
	Following bool
}
