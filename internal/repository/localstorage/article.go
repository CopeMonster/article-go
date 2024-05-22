package localstorage

import (
	"article-go/internal/domain"
	"context"
	"sync"
)

type LocalArticleRepository struct {
	articles map[int]*domain.Article
	mutex    *sync.Mutex
}

func NewLocalArticleRepository() (*LocalArticleRepository, error) {
	return &LocalArticleRepository{
		articles: make(map[int]*domain.Article),
		mutex:    new(sync.Mutex),
	}, nil
}

func (r *LocalArticleRepository) GetArticle(ctx context.Context) ([]*domain.Article, error) {
	r.mutex.Lock()

	defer r.mutex.Unlock()

	var articles []*domain.Article

	for _, v := range r.articles {
		articles = append(articles, v)
	}

	return articles, nil
}

func (r *LocalArticleRepository) GetArticleByID(ctx context.Context, id int) (*domain.Article, error) {
	r.mutex.Lock()

	defer r.mutex.Unlock()

	for k, v := range r.articles {
		if k == id {
			return v, nil
		}
	}

	return nil, domain.ErrArticleNotFound
}

func (r *LocalArticleRepository) CreateArticle(ctx context.Context, user *domain.User, article *domain.Article) error {
	r.mutex.Lock()
	r.articles[article.ID] = article
	r.mutex.Unlock()

	return nil
}

func (r *LocalArticleRepository) UpdateArticle(ctx context.Context, user *domain.User, id int, article domain.Article) error {
	r.mutex.Lock()

	defer r.mutex.Unlock()

	for _, v := range r.articles {
		if v.ID == id && v.AuthorID == user.ID {
			v.Title = article.Title
			v.Text = article.Text
			v.Views = article.Views
			v.UpdatedAt = article.UpdatedAt
			return nil
		}
	}

	return domain.ErrArticleNotFound
}

func (r *LocalArticleRepository) DeleteArticle(ctx context.Context, user *domain.User, id int) error {
	r.mutex.Lock()

	defer r.mutex.Unlock()

	for _, v := range r.articles {
		if v.ID == id && v.AuthorID == user.ID {
			delete(r.articles, id)
			return nil
		}
	}

	return domain.ErrArticleNotFound
}
