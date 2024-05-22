package repository

import (
	"article-go/internal/domain"
	"context"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, usernameOrEmail, password string) (*domain.User, error)
}

type ArticleRepository interface {
	GetArticle(ctx context.Context) ([]*domain.Article, error)
	GetArticleByID(ctx context.Context, id int) (*domain.Article, error)
	CreateArticle(ctx context.Context, user *domain.User, article *domain.Article) error
	UpdateArticle(ctx context.Context, user *domain.User, id int, article *domain.Article) error
	DeleteArticle(ctx context.Context, user *domain.User, id int) error
	UpdateArticleViews(ctx context.Context, article *domain.Article) error
}
