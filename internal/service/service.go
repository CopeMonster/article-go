package service

import (
	"article-go/internal/domain"
	"context"
)

const CtxUserKey = "user"

type AuthService interface {
	SignUp(ctx context.Context, username, email, password string) error
	SignIn(ctx context.Context, usernameOrEmail, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*domain.User, error)
}

type ArticleService interface {
	GetArticle(ctx context.Context) ([]*domain.Article, error)
	GetArticleByID(ctx context.Context, id int) (*domain.Article, error)
	CreateArticle(ctx context.Context, user *domain.User, title, text string) error
	UpdateArticle(ctx context.Context, user *domain.User, id int, title, text string) error
	DeleteArticle(ctx context.Context, user *domain.User, id int) error
}
