package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")

	ErrArticleNotFound = errors.New("article not found")
)
