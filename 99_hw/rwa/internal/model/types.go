package model

type Userr struct {
	DataUser `json:"user"`
}

// DataUser - структура, для хранения данных
type DataUser struct {
	ID       string `json:"id" testdiff:"ignore"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TestProfile struct {
	ID        string   `json:"id" testdiff:"ignore"`
	Email     string   `json:"email"`
	CreatedAt FakeTime `json:"createdAt"`
	UpdatedAt FakeTime `json:"updatedAt"`
	Username  string   `json:"username"`
	Bio       string   `json:"bio"`
	Image     string   `json:"image"`
	Token     string   `json:"token" testdiff:"ignore"`
	Following bool
}

type Response struct {
	User TestProfile `json:"User"`
}

type FakeTime struct {
	Valid bool `json:"Valid"`
}
