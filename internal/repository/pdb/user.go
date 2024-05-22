package pdb

import (
	"article-go/internal/domain"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresAuthRepository struct {
	db *sql.DB
}

func NewPostgresAuthRepository(connString string) (*PostgresAuthRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresAuthRepository{
		db: db,
	}, nil
}

func (r *PostgresAuthRepository) Init() error {
	return r.createTable()
}

func (r *PostgresAuthRepository) createTable() error {
	query := `create table if not exists users (
    	id serial primary key,
    	user_name varchar(50),
    	email varchar(50),
    	password varchar(50),
    	created_at timestamp
	)`

	_, err := r.db.Exec(query)

	return err
}

func (r *PostgresAuthRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `insert into users (user_name, email, password, created_at) values ($1, $2, $3, $4) RETURNING id`

	var newID int
	err := r.db.QueryRow(query,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt).Scan(&newID)

	if err != nil {
		return err
	}

	user.ID = newID

	return nil
}

func (r *PostgresAuthRepository) GetUser(ctx context.Context, usernameOrEmail, password string) (*domain.User, error) {
	query := `select * from users where (user_name=$1 or email=$2) and password=$3`

	rows, err := r.db.Query(query, usernameOrEmail, usernameOrEmail, password)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user, err := scanIntoUser(rows)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, err
}

func scanIntoUser(rows *sql.Rows) (*domain.User, error) {
	user := new(domain.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}
