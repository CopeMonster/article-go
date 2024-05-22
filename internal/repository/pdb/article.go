package pdb

import (
	"article-go/internal/domain"
	"context"
	"database/sql"
)

type PostgresArticleRepository struct {
	db *sql.DB
}

func NewPostgresArticleRepository(connString string) (*PostgresArticleRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &PostgresArticleRepository{
		db: db,
	}, nil
}

func (r *PostgresArticleRepository) Init() error {
	return r.createTable()
}

func (r *PostgresArticleRepository) createTable() error {
	query := `create table if not exists article (
    	id serial primary key,
    	author_id serial references users (id),
    	title varchar(50),
    	text text,
    	views serial,
    	created_at timestamp,
    	updated_at timestamp
	)`

	_, err := r.db.Exec(query)
	return err
}

func (r *PostgresArticleRepository) GetArticle(ctx context.Context) ([]*domain.Article, error) {
	query := `select * from article`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var articles []*domain.Article

	for rows.Next() {
		article, err := scanIntoArticle(rows)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (r *PostgresArticleRepository) GetArticleByID(ctx context.Context, id int) (*domain.Article, error) {
	query := `select * from article where id=$1`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoArticle(rows)
	}

	return nil, domain.ErrArticleNotFound
}

func (r *PostgresArticleRepository) CreateArticle(ctx context.Context, user *domain.User, article *domain.Article) error {
	query := `insert into article (
                     author_id, 
                     title, text, 
                     views, 
                     created_at, 
                     updated_at) values ($1, $2, $3, $4, $5, $6)`

	var newID int

	err := r.db.QueryRow(query,
		user.ID,
		article.Title,
		article.Text,
		article.Views,
		article.CreatedAt,
		article.UpdatedAt).Scan(&newID)

	article.ID = newID

	return err
}

func (r *PostgresArticleRepository) UpdateArticle(ctx context.Context, user *domain.User, id int, article *domain.Article) error {
	query := `update article set title=$1,
                   text = $2,
                   updated_at = $3
                   where id = $4 and author_id = $5`

	_, err := r.db.Exec(query,
		article.Title,
		article.Text,
		article.UpdatedAt,
		id,
		user.ID)

	return err
}

func (r *PostgresArticleRepository) DeleteArticle(ctx context.Context, user *domain.User, id int) error {
	query := `delete from article 
       where id = $1 and author_id = $2`

	_, err := r.db.Query(query, id, user.ID)

	return err
}

func (r *PostgresArticleRepository) UpdateArticleViews(ctx context.Context, article *domain.Article) error {
	query := `update article
               set views = $1
               where id = $2`

	_, err := r.db.Exec(query, article.Views, article.ID)

	return err
}

func scanIntoArticle(rows *sql.Rows) (*domain.Article, error) {
	article := new(domain.Article)

	err := rows.Scan(
		&article.ID,
		&article.AuthorID,
		&article.Title,
		&article.Text,
		&article.Views,
		&article.CreatedAt,
		&article.UpdatedAt)

	return article, err
}
