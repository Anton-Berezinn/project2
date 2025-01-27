package model

import "time"

type Userr struct {
	DataUser `json:"user"`
}

// DataUser - структура, для хранения данных
type DataUser struct {
	ID       string `json:"id" testdiff:"ignore"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type TestProfile struct {
	ID        string    `json:"id" testdiff:"ignore"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Image     string    `json:"image"`
	Token     string    `json:"token" testdiff:"ignore"`
	Following bool
}

type Response struct {
	User TestProfile `json:"User"`
}

type FakeTime struct {
	Valid bool `json:"Valid"`
}

type Article struct {
	TestArticle `json:"article"`
}

type TestArticle struct {
	Author         TestProfile `json:"author"`
	Body           string      `json:"body"`
	CreatedAt      time.Time   `json:"createdAt"`
	Description    string      `json:"description"`
	Favorited      bool        `json:"favorited"`
	FavoritesCount int         `json:"favoritesCount"`
	Slug           string      `json:"slug" testdiff:"ignore"`
	TagList        []string    `json:"tagList"`
	Title          string      `json:"title"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

type Artic struct {
	Article TestArticle `json:"article"`
}
