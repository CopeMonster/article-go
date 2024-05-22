package service

import (
	"article-go/internal/domain"
	"article-go/internal/repository"
	"context"
	"time"
)

type ArticleServiceImpl struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{
		repo: repo,
	}
}

func (s *ArticleServiceImpl) GetArticle(ctx context.Context) ([]*domain.Article, error) {
	return s.repo.GetArticle(ctx)
}

func (s *ArticleServiceImpl) GetArticleByID(ctx context.Context, id int) (*domain.Article, error) {
	article, err := s.repo.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	article.Views++

	err = s.repo.UpdateArticleViews(ctx, article)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleServiceImpl) CreateArticle(ctx context.Context, user *domain.User, title, text string) error {
	article := &domain.Article{
		AuthorID:  user.ID,
		Title:     title,
		Text:      text,
		Views:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateArticle(ctx, user, article)
}

func (s *ArticleServiceImpl) UpdateArticle(ctx context.Context, user *domain.User, id int, title, text string) error {
	article := &domain.Article{
		Title:     title,
		Text:      text,
		UpdatedAt: time.Now(),
	}

	return s.repo.UpdateArticle(ctx, user, id, article)
}

func (s *ArticleServiceImpl) DeleteArticle(ctx context.Context, user *domain.User, id int) error {
	return s.repo.DeleteArticle(ctx, user, id)
}
