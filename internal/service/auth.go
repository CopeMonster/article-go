package service

import (
	"article-go/internal/domain"
	"article-go/internal/repository"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

type AuthServiceImpl struct {
	repo           repository.AuthRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthService(
	repo repository.AuthRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) AuthService {

	return &AuthServiceImpl{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: tokenTTLSeconds * time.Second,
	}
}

func (s *AuthServiceImpl) SignUp(ctx context.Context, username, email, password string) error {
	pwd := sha1.New()

	pwd.Write([]byte(password))
	pwd.Write([]byte(s.hashSalt))

	user := &domain.User{
		Username:  username,
		Email:     email,
		Password:  fmt.Sprintf("%x", pwd.Sum(nil)),
		CreatedAt: time.Now(),
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *AuthServiceImpl) SignIn(ctx context.Context, usernameOrEmail, password string) (string, error) {
	pwd := sha1.New()

	pwd.Write([]byte(password))
	pwd.Write([]byte(s.hashSalt))

	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := s.repo.GetUser(ctx, usernameOrEmail, password)
	if err != nil {
		return "", domain.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(s.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.signingKey)
}

func (s *AuthServiceImpl) ParseToken(ctx context.Context, accessToken string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, domain.ErrInvalidAccessToken
}
