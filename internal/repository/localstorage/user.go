package localstorage

import (
	"article-go/internal/domain"
	"context"
	"sync"
)

type LocalAuthRepository struct {
	users map[int]*domain.User
	mutex *sync.Mutex
}

func NewLocalAuthRepository() (*LocalAuthRepository, error) {
	return &LocalAuthRepository{
		users: map[int]*domain.User{},
		mutex: new(sync.Mutex),
	}, nil
}

func (r *LocalAuthRepository) CreateUser(ctx context.Context, user *domain.User) error {
	r.mutex.Lock()
	r.users[user.ID] = user
	r.mutex.Unlock()

	return nil
}

func (r *LocalAuthRepository) GetUser(ctx context.Context, usernameOrEmail, password string) (*domain.User, error) {
	r.mutex.Lock()

	defer r.mutex.Unlock()

	for _, u := range r.users {
		if (u.Username == usernameOrEmail || u.Email == usernameOrEmail) && u.Password == password {
			return u, nil
		}
	}

	return nil, domain.ErrUserNotFound
}
