package domain

import "time"

type Article struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"authorID"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetArticleResponse struct {
	Article *Article `json:"article"`
}

type GetArticlesResponse struct {
	Articles []*Article `json:"articles"`
}

type CreateArticleInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type UpdateArticleInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
