package app

import (
	"article-go/internal/config"
	"article-go/internal/repository/pdb"
	"article-go/internal/server"
	"article-go/internal/service"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type App struct {
	AuthService    service.AuthService
	ArticleService service.ArticleService
}

func NewApp() *App {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load cfg; %s", err.Error())
	}

	connString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", cfg.DB.Username, cfg.DB.Database, cfg.DB.Password)

	userRepo, err := pdb.NewPostgresAuthRepository(connString)
	if err != nil {
		log.Fatalf("failed to create auth repository: %s", err.Error())
	}

	err = userRepo.Init()
	if err != nil {
		log.Fatalf("failed to init auth repository: %s", err.Error())
	}

	articleRepo, err := pdb.NewPostgresArticleRepository(connString)
	if err != nil {
		log.Fatalf("failed to create article repository: %s", err.Error())
	}

	err = articleRepo.Init()
	if err != nil {
		log.Fatalf("failed to init article repository: %s", err.Error())
	}

	return &App{
		AuthService: service.NewAuthService(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl")),
		ArticleService: service.NewArticleService(articleRepo),
	}
}

func (a *App) Run() {
	if err := server.Run(viper.GetString("server.port"), a.AuthService, a.ArticleService); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
